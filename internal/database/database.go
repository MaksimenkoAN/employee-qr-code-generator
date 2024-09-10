package database

import (
	"database/sql"
	"employee-qr-code-generator/internal/config"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

var db *sql.DB

type EmployeeInfo struct {
	Name        string
	WorkPhone   string
	MobilePhone string
	Email       string
	Address     string
}

// InitDB функция подключения к БД
func InitDB() error {
	dbConfig := config.AppConfig.Database
	connStr := fmt.Sprintf(
		"server=%s;port=%d;user id=%s;password=%s;database=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name,
	)

	var err error
	db, err = sql.Open("mssql", connStr)
	if err != nil {
		return err
	}
	return db.Ping()
}

// Добавьте функции для работы с базой данных здесь
// Например, получение информации о сотруднике по ID
func GetEmployeeInfo(employeeID string) (string, string, string, error) {
	// Пример получения данных о сотруднике
	// Подключение к базе данных и выполнение запроса
	return "John Doe", "+1234567890", "john.doe@example.com", nil
}

// GetUserId функция для получения ID ролей пользователя из БД
func GetUserId(username string) []int {
	var allUserIds []int

	query := `
        SELECT r.RoleId 
        FROM [BugNet_Employes].[dbo].[BugNet_UserRoles] r 
        WHERE r.UserId = (
            SELECT u.UserId 
            FROM [BugNet_Employes].[dbo].[Users] u 
            WHERE u.Username = ?
        )`
	rows, err := db.Query(query, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var userId int
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			panic(err)
		}
		allUserIds = append(allUserIds, userId)
	}

	return allUserIds
}

func GetEmployeeMobileFromZupp(employeeID string) string {
	var result sql.NullString

	query := `SELECT [MobilePhone] FROM [BugNet_Employes].[dbo].[Employes_ZUP] WHERE EmployeeID = ?`
	rows, err := db.Query(query, employeeID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			panic(err)
		}
	}
	return result.String
}

// GetInfoEmployee функция, которая по EmployeeID получает информация из БД. Если у пользователя есть привелегии, то также в результат будет добавлен мобильный номер телефона.
func GetInfoEmployee(employeeid string, privilegies bool) EmployeeInfo {
	var result EmployeeInfo
	query := `SELECT [EmployeeID], [Name], [MidName], [SurName], [Login], [Organization], [WorkPhone], [MobilePhone], [Chief], [Secretary], [Email], [Room], [LastLogOn], [DistinguishedName] FROM [BugNet_Employes].[dbo].[Employes_AD] WHERE EmployeeID = ?`
	rows, err := db.Query(query, employeeid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	//var str []string
	var empID, name, midName, surName, login, org, workPhone, mobilePhone, chief, secretary, email, room, lastLogOn, distinguishedName sql.NullString
	for rows.Next() {
		err := rows.Scan(&empID, &name, &midName, &surName, &login, &org, &workPhone, &mobilePhone, &chief, &secretary, &email, &room, &lastLogOn, &distinguishedName)
		if err != nil {
			panic(err)
		}
		result.Name = surName.String + " " + name.String + " " + midName.String
		result.WorkPhone = workPhone.String
		result.Address = room.String
		result.Email = email.String
		if privilegies {
			result.MobilePhone = GetEmployeeMobileFromZupp(employeeid)
		}
	}

	return result
}

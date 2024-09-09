package database

import (
	"database/sql"
	"employee-qr-code-generator/internal/config"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

var db *sql.DB

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

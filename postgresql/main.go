package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

func (e *Employee) Print() {
	fmt.Println("ID :", e.ID)
	fmt.Println("FullName :", e.FullName)
	fmt.Println("Email :", e.Email)
	fmt.Println("Age :", e.Age)
	fmt.Println("Division :", e.Division)
	fmt.Println()
}

const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_NAME = "db-go-sql"
)

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	// UNCOMMENT CODE BELOW TO TRY

	// create employee
	// emp := Employee{
	// 	Email:    "admin@noobeeid.com",
	// 	FullName: "Noobeeid",
	// 	Age:      22,
	// 	Division: "Project Manager",
	// }

	// err = createEmployee(db, &emp)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	employees, err := getAllEmployees(db)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for _, emp := range *employees {
		emp.Print()
	}

	// fmt.Println("====== Get Employee by id ======")
	// employee, err := getEmployeeById(db, 2)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// employee.Print()

	// Update Employee
	// newEmp := Employee{
	// 	Email:    "update@mail.com",
	// 	FullName: "Test Update",
	// 	Age:      21,
	// 	Division: "QA",
	// }

	// err = updateEmployee(db, &newEmp, 2)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// Delete Employee
	// err = DeleteEmployee(db, 2)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)

	return db, nil
}

func getAllEmployees(db *sql.DB) (*[]Employee, error) {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var employees []Employee

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division)

		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return &employees, nil
}

func createEmployee(db *sql.DB, request *Employee) error {
	query := `
		INSERT INTO employees(full_name, email, age, division)
		VALUES ($1, $2, $3, $4)
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.FullName, request.Email, request.Age, request.Division)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func getEmployeeById(db *sql.DB, id int) (*Employee, error) {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
		WHERE id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)

	var emp Employee
	err = row.Scan(&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func updateEmployee(db *sql.DB, r *Employee, id int) error {
	query := `
		UPDATE employees
		SET full_name = $2, email = $3, age = $4, division = $5
		WHERE id = $1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(query, id, r.FullName, r.Email, r.Age, r.Division)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("Updated data amount :", count)
	return nil
}

func DeleteEmployee(db *sql.DB, id int) error {
	query := `
		DELETE FROM employees
		WHERE id = $1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Deleted data amount :", count)
	return nil
}

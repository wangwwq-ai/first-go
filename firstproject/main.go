package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	SID     int
	SName   string
	SAge    int
	ClassID int
}

func main() {
	// 数据库连接信息
	db, err := sql.Open("mysql", "root:wwq123@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 查询所有学生数据
	rows, err := db.Query("SELECT sid, sname, sage, classid FROM student")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.SID, &student.SName, &student.SAge, &student.ClassID)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}

	// 检查是否有错误发生
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// HTML模板
	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Student Information</title>
    <style>
        table {
            width: 50%;
            border-collapse: collapse;
            margin: 20px auto;
        }
        th, td {
            border: 1px solid #ddd;
            padding: 8px;
            text-align: center;
        }
        th {
            background-color: #f2f2f2;
        }
    </style>
</head>
<body>
    <h2 style="text-align:center;">Student Information</h2>
    <table>
        <thead>
            <tr>
                <th>SID</th>
                <th>Name</th>
                <th>Age</th>
                <th>Class ID</th>
            </tr>
        </thead>
        <tbody>
            {{range .}}
            <tr>
                <td>{{.SID}}</td>
                <td>{{.SName}}</td>
                <td>{{.SAge}}</td>
                <td>{{.ClassID}}</td>
            </tr>
            {{end}}
        </tbody>
    </table>
</body>
</html>
`

	t := template.Must(template.New("webpage").Parse(tmpl))
	f, err := os.Create("students.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = t.Execute(f, students)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully generated students.html")
}

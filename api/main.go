package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"com.ikhsanhaikal/backendtest/digital-nayaka-abhinaya/api/sqlcdb"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type ResponseSoal5 struct {
	Employee           sqlcdb.Employee
	Perkiraan_kenaikan float64
	TotalUlasan        int
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed on load env file")
	}

	ctx := context.Background()
	url := fmt.Sprintf("postgres://postgres:%s@localhost:55000/%s", os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	fmt.Println(url)
	conn, err := pgx.Connect(ctx, url)

	if err != nil {
		log.Fatal("failed on connecting db, err: ", err.Error())
	}

	defer conn.Close(ctx)

	app := fiber.New()

	app.Get("/soal2", func(c *fiber.Ctx) error {
		queries := sqlcdb.New(conn)

		employess, err := queries.SoalDua(c.Context())

		if err != nil {
			c.SendStatus(http.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"employees": employess,
		})
	})

	app.Get("/soal3", func(c *fiber.Ctx) error {
		queries := sqlcdb.New(conn)

		employess, err := queries.SoalTiga(c.Context())

		if err != nil {
			c.SendStatus(http.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"employees": employess,
		})
	})

	app.Get("/soal4", func(c *fiber.Ctx) error {
		queries := sqlcdb.New(conn)

		result, err := queries.SoalEmpat(c.Context())

		if err != nil {
			c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"result": result,
		})
	})

	app.Get("/soal5", func(c *fiber.Ctx) error {
		queries := sqlcdb.New(conn)

		totalUlasanPerId, err := queries.SoalLimaCountTotalUlasan(c.Context())
		fmt.Printf("totalUlasanperId: %+v\n", totalUlasanPerId)
		if err != nil {
			c.SendStatus(http.StatusInternalServerError)
		}
		employees, err := queries.SoalLima(c.Context())
		if err != nil {
			c.SendStatus(http.StatusInternalServerError)
		}

		// fmt.Printf("emp: %+v\n", employees)

		result := []ResponseSoal5{}
		for _, emp := range employees {
			fmt.Printf("emp: %+v\n", emp)
			converted, _ := emp.Salary.Float64Value()
			perkiraan_kenaikan := converted.Float64
			for i := 0; i < 7; i++ { // 2009-2016
				perkiraan_kenaikan += perkiraan_kenaikan * 15 / 100
			}

			totalUlasan := 0
			for _, val := range totalUlasanPerId {
				if val.ID == emp.ID {
					totalUlasan = int(val.Total)
				}
			}

			result = append(result, ResponseSoal5{
				Employee:           emp,
				Perkiraan_kenaikan: perkiraan_kenaikan,
				TotalUlasan:        totalUlasan,
			})
		}

		return c.JSON(fiber.Map{
			"result": result,
		})
	})

	app.Listen("localhost:3000")

}

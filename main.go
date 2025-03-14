package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"net/url"
	"os"
	"time"
	
	
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/lengzuo/supa"
)

type Quotes struct {
	No        int    `json:"no"`
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}



func random() int {
	return rand.Intn(8611)
}

func connectWithSupa() (error,*supabase.Client) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error loading env ")
	}
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")
	conf := supabase.Config{
		ApiKey:     key,
		ProjectRef: url,
		Debug:      true,
	}
	supaClient, err := supabase.New(conf)
	if err != nil {
		log.Println("failed to init supa Client:", err)
		
	}
	return err,supaClient
}

func getData(c *fiber.Ctx) error {
	k := random()
	err,supaClient := connectWithSupa()
	ctx := context.Background()
	var u []Quotes
	query := supaClient.DB.From("anime_quotes").Select("*").Eq("No", strconv.Itoa(k) )
	err = query.Execute(ctx, &u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to execute query")
	}
	bytes , err := json.Marshal(u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed in Marshal")
	}
	log.Println(bytes)
	return c.JSON(u)
}

func dataById(c *fiber.Ctx) error {
	err,supaClient := connectWithSupa()
	ctx := context.Background()
	id := c.Params("id")
	var u []Quotes
	query := supaClient.DB.From("anime_quotes").Select("*").Eq("No",id )
	err = query.Execute(ctx, &u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to execute query")
	}
	bytes , err := json.Marshal(u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed in Marshal")
	}
	log.Println(bytes)
	return c.JSON(u)
}

func dataByAnimeName(c *fiber.Ctx) error {
	err,supaClient := connectWithSupa()
	ctx := context.Background()
	name := c.Params("name")
	
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid URL parameter")
	}
	formattedName := "(" + decodedName + ")"
	
	var u []Quotes
	query := supaClient.DB.From("anime_quotes").Select("*").Eq("Anime",formattedName )
	
	err = query.Execute(ctx, &u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to execute query")
	}
	bytes , err := json.Marshal(u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed in Marshal")
	}
	log.Println(bytes)
	return c.JSON(u)
}

func dataByCharacterName(c *fiber.Ctx) error {
	err,supaClient := connectWithSupa()
	ctx := context.Background()
	name := c.Params("name")
	
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid URL parameter")
	}

	
	var u []Quotes
	query := supaClient.DB.From("anime_quotes").Select("*").Eq("Character",decodedName )
	log.Printf("Constructed query: %v", query)
	err = query.Execute(ctx, &u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to execute query")
	}
	bytes , err := json.Marshal(u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed in Marshal")
	}
	log.Println(bytes)
	return c.JSON(u)
}


func main(){
	app := fiber.New()
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        10,
	}))
	app.Get("/quote",getData)
	app.Get("/quote/:id", dataById)
	app.Get("/quote/anime/:name", dataByAnimeName)
	app.Get("/quote/character/:name", dataByCharacterName)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen("0.0.0.0:" + port)
}
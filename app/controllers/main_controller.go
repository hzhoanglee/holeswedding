package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"hz-evite/app/models"
	"hz-evite/pkg/database"
	"image"
	_ "image/jpeg"
	"os"
)

type Image struct {
	ImagePath string `json:"image_path"`
	Title     string `json:"title"`
	Quote     string `json:"quote"`
	Size      string `json:"size"`
}

func RenderHello(c *fiber.Ctx) error {
	guestId := c.Params("guestId")
	if guestId == "" {
		guestId = "00000000-0000-0000-0000-000000000000"
	}
	user := models.User{}
	db := database.DB
	if err := db.First(&user, "id = ?", guestId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	imageMap := make(map[string]Image)
	file, err := os.Open("image_map.json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error opening image map file")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&imageMap); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error decoding image map file")
	}

	return c.Render("index", fiber.Map{
		"guestName": user.Name,
		"imageMap":  imageMap,
	})
}

func oldImgMap() {
	imageMap := map[int]Image{
		1:  {ImagePath: "/anhcuoi/1.jpg", Title: "Forever Begins Today", Quote: "When you realize you want to spend the rest of your life with somebody, you want the rest of your life to start as soon as possible. - When Harry Met Sally"},
		2:  {ImagePath: "/anhcuoi/2.jpg", Title: "Two Hearts, One Love", Quote: "I have found the one whom my soul loves. - Song of Solomon 3:4"},
		3:  {ImagePath: "/anhcuoi/3.jpg", Title: "The Beginning of Always", Quote: "Whatever our souls are made of, his and mine are the same. - Emily Brontë, Wuthering Heights"},
		4:  {ImagePath: "/anhcuoi/4.jpg", Title: "Love's Sweet Promise", Quote: "In all the world, there is no heart for me like yours. In all the world, there is no love for you like mine. - Maya Angelou"},
		5:  {ImagePath: "/anhcuoi/5.jpg", Title: "Eternal Embrace", Quote: "I would rather share one lifetime with you than face all the ages of this world alone. - Lord of the Rings"},
		6:  {ImagePath: "/anhcuoi/6.jpg", Title: "Our Perfect Day", Quote: "You are my today and all of my tomorrows. - Leo Christopher"},
		7:  {ImagePath: "/anhcuoi/7.jpg", Title: "Whispers of Forever", Quote: "I love you not only for what you are, but for what I am when I am with you. - Elizabeth Barrett Browning"},
		8:  {ImagePath: "/anhcuoi/8.jpg", Title: "Dance of Two Souls", Quote: "We loved with a love that was more than love. - Edgar Allan Poe"},
		9:  {ImagePath: "/anhcuoi/9.jpg", Title: "Heaven on Earth", Quote: "To me, you are perfect. - Love Actually"},
		10: {ImagePath: "/anhcuoi/10.jpg", Title: "The Vow of Hearts", Quote: "Grow old along with me; the best is yet to be. - Robert Browning"},
		11: {ImagePath: "/anhcuoi/11.jpg", Title: "Infinite Devotion", Quote: "I choose you. And I'll choose you over and over and over. Without pause, without a doubt, in a heartbeat. I'll keep choosing you. - Unknown"},
		12: {ImagePath: "/anhcuoi/12.jpg", Title: "Love's Symphony", Quote: "Being deeply loved by someone gives you strength, while loving someone deeply gives you courage. - Lao Tzu"},
		13: {ImagePath: "/anhcuoi/13.jpg", Title: "United as One", Quote: "The best thing to hold onto in life is each other. - Audrey Hepburn"},
		14: {ImagePath: "/anhcuoi/14.jpg", Title: "A Dream Come True", Quote: "I knew the second I met you that there was something about you I needed. Turns out it wasn't something about you at all. It was just you. - Beautiful Disaster"},
		15: {ImagePath: "/anhcuoi/15.jpg", Title: "Written in the Stars", Quote: "When I saw you I fell in love, and you smiled because you knew. - William Shakespeare"},
		16: {ImagePath: "/anhcuoi/16.jpg", Title: "The Magic Moment", Quote: "You know you're in love when you can't fall asleep because reality is finally better than your dreams. - Dr. Seuss"},
		17: {ImagePath: "/anhcuoi/17.jpg", Title: "Boundless Love", Quote: "I love you without knowing how, or when, or from where. I love you simply, without problems or pride. - Pablo Neruda"},
		18: {ImagePath: "/anhcuoi/18.jpg", Title: "Heart's Desire", Quote: "If I had a flower for every time I thought of you... I could walk through my garden forever. - Alfred Tennyson"},
		19: {ImagePath: "/anhcuoi/19.jpg", Title: "Our Love Story", Quote: "Every love story is beautiful, but ours is my favorite. - Unknown"},
		20: {ImagePath: "/anhcuoi/20.jpg", Title: "Soulmate's Journey", Quote: "In a sea of people, my eyes will always search for you. - Unknown"},
		21: {ImagePath: "/anhcuoi/21.jpg", Title: "Destiny's Design", Quote: "I saw that you were perfect, and so I loved you. Then I saw that you were not perfect and I loved you even more. - Angelita Lim"},
		22: {ImagePath: "/anhcuoi/22.jpg", Title: "Cherished Moments", Quote: "A successful marriage requires falling in love many times, always with the same person. - Mignon McLaughlin"},
		23: {ImagePath: "/anhcuoi/23.jpg", Title: "Love's Pure Light", Quote: "You are the finest, loveliest, tenderest, and most beautiful person I have ever known—and even that is an understatement. - F. Scott Fitzgerald"},
		24: {ImagePath: "/anhcuoi/24.jpg", Title: "Together Forever", Quote: "I want all of you, forever, you and me, every day. - The Notebook"},
		25: {ImagePath: "/anhcuoi/25.jpg", Title: "The Promise Ring", Quote: "Love is composed of a single soul inhabiting two bodies. - Aristotle"},
		26: {ImagePath: "/anhcuoi/26.jpg", Title: "Endless Romance", Quote: "Once upon a time you were a wish I made on a shooting star. - Unknown"},
		27: {ImagePath: "/anhcuoi/27.jpg", Title: "Hearts Entwined", Quote: "I love you more than I have ever found a way to say to you. - Ben Folds"},
		28: {ImagePath: "/anhcuoi/28.jpg", Title: "Forever Yours", Quote: "My heart is and always will be yours. - Jane Austen, Sense and Sensibility"},
		29: {ImagePath: "/anhcuoi/29.jpg", Title: "Till Death Do Us Part", Quote: "I vow to fiercely love you in all your forms, now and forever. I promise to never forget that this is a once in a lifetime love. - The Vow"},
	}

	rootFolder := "assets"
	for i := 1; i <= len(imageMap); i++ {
		imagePath := rootFolder + imageMap[i].ImagePath

		// Open the image file
		file, err := os.Open(imagePath)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", imagePath, err)
			continue
		}
		defer file.Close()

		// Decode image to get dimensions
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Printf("Error decoding image %s: %v\n", imagePath, err)
			continue
		}

		// Get image bounds
		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()

		// Update the map with size information
		imgData := imageMap[i]
		imgData.Size = fmt.Sprintf("%dx%d", width, height)
		imageMap[i] = imgData

	}

	file, err := os.Create("image_map.json")
	if err != nil {

	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(imageMap); err != nil {

	}
}

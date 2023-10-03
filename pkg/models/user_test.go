package models_test

import (
	"github.com/ocrosby/go-mongo-tutorial/pkg/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Describe("Slug", func() {
		It("returns an empty string for the default user", func() {
			// Arrange
			user := models.NewUser()

			// Act
			slug := user.Slug()

			// Assert
			Expect(slug).To(Equal(""))
		})

		It("returns a slug for a user with a first and last name", func() {
			// Arrange
			user := models.NewUser()
			user.FirstName = "John"
			user.LastName = "Doe"

			// Act
			slug := user.Slug()

			// Assert
			Expect(slug).To(Equal("john-doe"))
		})
	})

	Describe("String", func() {
		It("returns a string representation of the user", func() {
			// Arrange
			user := models.NewUser()
			user.FirstName = "John"
			user.LastName = "Doe"

			// Act
			str := user.String()

			// Assert
			Expect(str).To(Equal("User: John Doe"))
		})

		It("returns a string representation of the user without a first name", func() {
			// Arrange
			user := models.NewUser()
			user.LastName = "Doe"

			// Act
			str := user.String()

			// Assert
			Expect(str).To(Equal("User:  Doe"))
		})

		It("returns a string representation of the user without a last name", func() {
			// Arrange
			user := models.NewUser()
			user.FirstName = "John"

			// Act
			str := user.String()

			// Assert
			Expect(str).To(Equal("User: John "))
		})

		It("returns a string representation of the user without a first or last name", func() {
			// Arrange
			user := models.NewUser()

			// Act
			str := user.String()

			// Assert
			Expect(str).To(Equal("User:  "))
		})
	})
})

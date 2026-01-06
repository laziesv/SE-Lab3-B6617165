package	entity

import (

	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestStudentValidation(t *testing.T) {
	govalidator.SetFieldsRequiredByDefault(false)
	RegisterTestingT(t)

	t.Run("Test Student Valid Case", func(t *testing.T){
		s := Student{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Age:       20,
		}
		ok, err := govalidator.ValidateStruct(s)
		Expect(ok).To(BeTrue())
		Expect(err).To(BeNil())
	})

	t.Run("Test Student Invalid Case - Missing FirstName", func(t *testing.T){
		s := Student{
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Age:       20,
		}
		ok, err := govalidator.ValidateStruct(s)
		Expect(ok).To(BeFalse())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("First name is required"))
	})

	t.Run("Test Student Invalid Case - Age Out of Range", func(t *testing.T){
		s := Student{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Age:       1000,
		}
		ok, err := govalidator.ValidateStruct(s)
		Expect(ok).To(BeFalse())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("Age must be between 1 and 150"))
	})

	t.Run("Test Student Invalid Case - Missing Email", func(t *testing.T){
		s := Student{
			FirstName: "John",
			LastName:  "Doe",
			Age:       20,
		}
		ok, err := govalidator.ValidateStruct(s)
		Expect(ok).To(BeFalse())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("Email is required"))
	})

	t.Run("Test Student Invalid Case - Missing LastName", func(t *testing.T){
		s := Student{
			FirstName: "John",
			Email:     "john.doe@example.com",
			Age:       20,
		}
		ok, err := govalidator.ValidateStruct(s)
		Expect(ok).To(BeFalse())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("Last name is required"))
	})
}
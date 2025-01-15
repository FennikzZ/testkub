package unit
import(
	"T2/entity"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"testing"
	"fmt"
)
func TestPhoneNumber(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Check Phone`,func(t *testing.T){
		user := entity.User{
			Phone: "",
			Email:		"F@gmailc.com",
		}
		ok,err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone is required"))
	})

	t.Run(`phone_number`, func(t *testing.T) {
		user := entity.User{
			Phone:     "080800000000", // ผิดตรงนี้ มี 11 ตัว
			Email:		"F@gmailc.com",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Phone: %s does not validate as stringlength(10|10)", user.Phone)))

	})
}

func TestEmailKub(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Check Email Kub`,func(t *testing.T){
		user := entity.User{
			Phone:	"0988675090",
			Email:	"",
		}
		ok,err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("email is required"))
	})

	t.Run(`Check Email Ja`,func(t *testing.T){
		user := entity.User{
			Phone:	"0988675090",
			Email:	"123@@.como",
		}
		ok,err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})
}
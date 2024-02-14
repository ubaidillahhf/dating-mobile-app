package helper

import (
	"fmt"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func RandomUsername(fullname string) string {
	fullnameLowercase := strings.ToLower(fullname)
	firstName := strings.Split(fullnameLowercase, " ")[0]
	shortIDLowercase := strings.ToLower(gonanoid.Must())

	return fmt.Sprintf("%s%s", firstName, strings.Replace(shortIDLowercase, "-", "", -1))
}

// input example: Porto.Feature[0].FeatureList
// output example: Porto.Feature[0].featureList
func ConvLastStructNameToCamelCase(data string) string {

	strArr := strings.Split(data, ".")

	lastIndex := len(strArr) - 1
	lastStr := strArr[lastIndex]

	convLastStr := strcase.ToLowerCamel(lastStr)
	strArr[lastIndex] = convLastStr

	return strings.Join(strArr[1:], ".")

}

func GetAgeFromDob(dob time.Time) string {

	dobYear := dob.Year()
	if dobYear == 1 {
		return "undisclosed"
	}

	now := time.Now()
	age := now.Year() - dobYear

	if now.YearDay() < dob.YearDay() {
		age--
	}

	return fmt.Sprintf("%d yo", age)

}

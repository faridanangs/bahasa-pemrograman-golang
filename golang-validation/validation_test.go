package golangvalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	name := "d"
	// kita melakukan validasi terhadap variable name jika tidak ada nilai dan < 1 nilai di dalamnya maka dia akan terjadi error
	err := validate.Var(name, "required,min=2")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationVariables(t *testing.T) {
	validate := validator.New()

	name1 := "farid"
	name2 := "farid"

	// supaya kita bisa memasukan dua variable sekaligun kita menggunakan VarWithValue
	// kemudian kita melakukan pengecekan terhadap dua variable menggunakan eqfield jika name1 == name2 maka dia tidak error namun
	// jika dia berbeda maka dia akan error
	err := validate.VarWithValue(name1, name2, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationmultipleTag(t *testing.T) {
	validate := validator.New()

	number := "farid"
	// untuk melakukan multiple tag kita bisa menggunakan , sebagai pemsahnya
	err := validate.Var(number, "required,numeric")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationParameterTag(t *testing.T) {
	validate := validator.New()
	number := "1255"
	// untuk melakukan parameter tag kita bisa menggunakan = untuk memberi nilai yang harus di penuhi
	err := validate.Var(number, "required,numeric,min=3,max=6")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// untuk melakukan struct validation kita bisa menambahan `validate:"tagnta" json:"username"`
type LoginUser struct {
	Username string `validate:"required,email"`
	Password string `validate:"required,min=5"`
}

func TestValidationStruct(t *testing.T) {
	validate := validator.New()
	loginUser := &LoginUser{
		Username: "farid@example.com",
		Password: "farid1121",
	}
	// kemudian untuk memvalidasi structnya kita bisa menggunakan validate.Struct()
	err := validate.Struct(loginUser)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	validate := validator.New()
	loginUser := &LoginUser{
		Username: "farid",
		Password: "fari",
	}
	// kemudian untuk memvalidasi structnya kita bisa menggunakan validate.Struct()
	err := validate.Struct(loginUser)
	if err != nil {

		// untuk melakukan error  validate kita konversi nilai err menjadi ValidationErrors.
		validationError := err.(validator.ValidationErrors)
		// kemudian kita itrasi semua data errornya dan kita ambil data field error yang ingin kita tampilkan
		for _, fieldError := range validationError {
			fmt.Println("error: ", fieldError.Field(), "on tag: ", fieldError.Tag(), "errors: ", fieldError.Error())
		}
	}
}

func TestValidationStructField(t *testing.T) {
	type LoginUser struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
		// kita menambahkan eqfield di dalam ConfirmPassword di mana ini bertujuan untuk melakukan
		// validation apakah ConfirmPassword == password jika sama dia tidk error
		ConfirmPassword string `validate:"required,min=5,nefield=Password"`
	}
	validate := validator.New()
	loginUser := &LoginUser{
		Username:        "farid@exa.com",
		Password:        "faridanangs",
		ConfirmPassword: "faridanangs",
	}
	// kemudian untuk memvalidasi structnya kita bisa menggunakan validate.Struct()
	err := validate.Struct(loginUser)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationStructNested(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}
	// kita tidak perlu hawatir jika kita memiliki struct bersarang valiadtor tetap akan
	// memvalidatenya tanpa terkecuali
	type User struct {
		Id      int     `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}
	validate := validator.New()
	request := User{
		Id:   9,
		Name: "farid",
		Address: Address{
			City:    "",
			Country: "",
		},
	}
	// kemudian untuk memvalidasi structnya kita bisa menggunakan validate.Struct()
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationColection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}
	type User struct {
		Id   string `validate:"required"`
		Name string `validate:"required"`
		// supaya validate di dalam slice ini bisa di baca dan di validate kita tambah tag dive di dalam addressnya
		Address []Address `validate:"required,dive"`
	}
	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Address: []Address{
			{City: ""},
			{Country: ""},
		},
	}
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// kita membuat slice bertipe string dll
func TestValidationBasicColection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}
	type User struct {
		Id      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`

		// di sini kita membuat Hobies ini required dan kita menggunakan dive untk masuk ke dalam
		// slice dan memberi tag required dan min=3 untuk semua data slice di dalamnya
		Hobies []string `validate:"required,dive,required,min=3"`
	}
	validate := validator.New()
	request := User{
		Id:   "389083092",
		Name: "farid anang s",
		Address: []Address{
			{City: "mataram"},
			{Country: "lombok"},
		},
		Hobies: []string{
			"Gaming", "Coding", "makan",
		},
	}
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}
	type School struct {
		Name string `validate:"required"`
	}
	type User struct {
		Id      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
		Hobies  []string  `validate:"required,dive,required,min=3"`
		// kita masuk ke dalam map dan kita menyetel nilai dari keys harus di isi dengan maksimal
		// nilai 3 kemudian kita tutup keysnya dan masuk ke dalam value school menggunakan dive
		Schools map[string]School `validate:"dive,keys,required,min=3,endkeys,dive"`
	}
	validate := validator.New()
	request := User{
		Id:   "2091902",
		Name: "farid anang s",
		Address: []Address{
			{City: "sesela"},
			{Country: "lombok"},
		},
		Hobies: []string{
			"Gaming", "Coding", "makan", "tidur",
		},
		Schools: map[string]School{
			"SD": {
				Name: "min 2 lobar",
			},
			"SMP": {
				Name: "farid",
			},
			"SMA": {
				Name: "sman 1 gunungsari",
			},
		},
	}
	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMap(t *testing.T) {
	type Nama struct {
		Name string `validate:"required"`
	}
	type Hewans struct {
		Hewan map[string]Nama `validate:"dive,keys,required,min=3,endkeys,dive"`
	}

	dataHewan := Hewans{
		Hewan: map[string]Nama{
			"": {
				Name: "farid",
			},
		},
	}

	validate := validator.New()

	for key, value := range dataHewan.Hewan {
		if err := validate.Struct(value); err != nil {
			fmt.Printf("Error validating element with key %s: %s\n", key, err.Error())
		}
	}
}

func TestValidatioBasicnMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}
	type School struct {
		Name string `validate:"required"`
	}
	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobies    []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"required,dive,keys,required,min=2,endkeys,required"`
		Wallet    map[string]int    `validate:"required,dive,keys,required,endkeys,required,gt=10000"`
	}
	validate := validator.New()
	request := User{
		Id:   "09u0u3",
		Name: "farid anang s",
		Addresses: []Address{
			{
				City:    "mataram",
				Country: "indonesia"},
		},
		Hobies: []string{
			"Gaming", "Coding", "makan", "makan",
		},
		Schools: map[string]School{
			"sd": {
				Name: "min model",
			},
			"SMP": {
				Name: "min model",
			},
			"sma": {
				Name: "smansa",
			},
		},
		Wallet: map[string]int{
			"BRI":     10000,
			"":        100000000,
			"MANDIRI": 0,
		},
	}
	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationAlias(t *testing.T) {
	validate := validator.New()

	// kita membuat alias supaya validate ini bisa di gunakan di semua validate
	// tanpa harus menulisnya secara ulang cukup kta buat sekali dan kita panggil
	// berkali kali
	validate.RegisterAlias("varchar", "required,max=100")

	type Seller struct {
		// untuk cara memanggil nya kita cukup ketikan nama dari alias nya di dalam validate
		// kita juga bisa menambah tag lagi sesua dengan yang kita butuhkan
		Id     string `validate:"varchar,min=4"`
		Name   string `validate:"varchar"`
		Slogan string `validate:"varchar"`
		Owner  string `validate:"varchar"`
	}

	data := Seller{
		Id:     "313",
		Owner:  "",
		Slogan: "",
		Name:   "",
	}

	err := validate.Struct(data)
	if err != nil {
		fmt.Println(err)
	}
}

// kita membuat sebuah validator menggunakan validator.FieldLevel yang bertipe interface di dalam pkg validator
// dan harus menerima nilai bertipe FieldLevel
func CustomValdation(field validator.FieldLevel) bool {
	// kita mengambil nilai yang di kirim di validate di dalam field dan mengkonversinya ke dalam string
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()

	// kita registrasikan dulu nama dari tag yang akan kita buat dan func tempat kita akan jalankan data dari tag yang kita buat
	validate.RegisterValidation("username", CustomValdation)

	type User struct {
		Username string `validate:"username,max=7"`
		Password string `validate:"username"`
	}

	data := User{
		Username: "FAR",
		Password: "",
	}

	err := validate.Struct(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNum = regexp.MustCompile("^[0-9]+$")

func MustValidationPin(field validator.FieldLevel) bool {
	// unutk menangkap nilai dari parametrer kita busa menggunakn Param()
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNum.MatchString(value) {
		return false
	}

	return len(value) == length

}

func TestCustomValidationParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidationPin)

	type Handphon struct {
		Password string `validate:"required,number"`
		// unutk mengirim nilai ke dalam parameter kita taruh di dalam =
		Pin string `validate:"required,pin=6"`
	}

	data := Handphon{
		Password: "farid",
		Pin:      "123456",
	}

	err := validate.Struct(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationOr(t *testing.T) {
	validate := validator.New()

	type Handphon struct {
		// kita bisa menaruh or di dalam tag supaya dia bisa memilih salah satu dari tagnya
		// misal disini kita menaruh or di dalam tag username email or numeric. dia bisa memilih salah satunya
		Username string `validate:"required,email|numeric"`
		Pin      string `validate:"required"`
	}

	data := Handphon{
		Username: "farid@exe.com",
		Pin:      "123",
	}

	err := validate.Struct(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustEqIgnorCase(field validator.FieldLevel) bool {
	//// jika menggunakan Param kita mengambil nilai dari parameter yang di kirim
	// contoh username=Password yang di ambil string Passwordnya budan data yang ada di dalam password
	// firstValue := strings.ToUpper(field.Field().String())
	// secondValue := strings.ToUpper(field.Param())
	// return firstValue == secondValue

	// oleh karna itu untuk menangkap data di salam Password/Email kita menggunakan
	// interface dari GetStructFieldOK2

	// value,typeData,nil|not,ok|not := ...
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("Not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	fmt.Println("value.String: ", value.String())
	fmt.Println("field.Field: ", field.Field().String())

	return firstValue == secondValue

}

func TestCrossValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equal", MustEqIgnorCase)

	type User struct {
		Username string `validate:"required,field_equal=Password|field_equal=Email"`
		Password string `validate:"required,number"`
		Email    string `validate:"required,email"`
		Name     string `validate:"required"`
	}
	data := User{
		Username: "1309128390",
		Password: "1309128390",
		Email:    "farid@exa.com",
		Name:     "farid",
	}

	err := validate.Struct(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}
type RegisterR struct {
	User  string `validate:"required"`
	Ema   string `validate:"required,email"`
	Pho   string `validate:"required,numeric"`
	Passw string `validate:"required"`
}

// kita meakukan validation terhadap struct menggunakan interface StructLevel
func MustValidRegisterSucces(field validator.StructLevel) {
	// untuk mendapatkan data di dalamnya kita gunakan  field.Current().Interface().(kita korversi menjadi>RegisterRequest)
	// data yang kita konversi harus sama dengan data yang kita kirim di RegisterStructValidation()
	request := field.Current().Interface().(RegisterRequest)

	if request.Username == request.Email || request.Username == request.Phone {
		// success
	} else {
		// jika gagal kita mengirimkan error
		field.ReportError(request.Username, "Username", "Username", "'username", "error sayang")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validate := validator.New()

	// kita registrasi dulu dan mengirim func MustValidRegisterSucces dan struct RegisterRequest
	// data yang kita konversi harus sama dengan data yang kita kirim biar tidak terjadi maslah
	validate.RegisterStructValidation(MustValidRegisterSucces, RegisterRequest{})

	data := RegisterRequest{
		Username: "faid@exa.com",
		Email:    "farid@exa.com",
		Phone:    "698649249",
		Password: "farid",
	}

	err := validate.Struct(data)

	if err != nil {
		fmt.Println(err.Error())
	}
}


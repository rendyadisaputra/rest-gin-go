package miscs
import (
	"fmt"
	"os"
	"github.com/joho/godotenv"

)
var err = godotenv.Load()
var ReleaseMode = os.Getenv("RELEASE_MODE");
func Var_dump(expression ...interface{} ) {
	
	if(ReleaseMode!="1"){
		fmt.Println(fmt.Sprintf("%#v", expression))
	}
}

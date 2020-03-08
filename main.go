package leads

import "github.com/gin-gonic/gin"

func main() {
	routr := core.CreateContext()
	defer core.Shutdown()

	routr
}

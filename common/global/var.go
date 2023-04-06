// description:
// @author renshiwei
// Date: 2022/10/5 17:06

package global

var (
	Config ConfigYaml
)

type ConfigYaml struct {
	Server struct {
		Name    string
		Version string
		Http    struct {
			Host string
			Port int
		}
	}

	Cli struct {
		Name string
	}

	Log struct {
		Level struct {
			Server string
		}
	}

	Eth struct {
		ElAddr        string `json:"elAddr"`
		ElArchiveAddr string `json:"elArchiveAddr"`
		ClAddr        string `json:"clAddr"`
	} `json:"eth"`
}

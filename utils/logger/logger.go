package logger

import (
	"github.com/cihub/seelog"
	"log"
)

func init() {
	config :=  `
<seelog type="asynctimer" asyncinterval="1000000" minlevel="debug" maxlevel="error">
	<outputs formatid="kube">
		<console/>
		<rollingfile formatid="kube" type="size" filename="/var/log/mixer.log" maxsize="1000000" maxrolls="5" />
	</outputs>
	<formats>
		<format id="kube" format="[%Date %Time %LEVEL %FullPath %Func %Line]: %Msg%n"/>
	</formats>
</seelog>`
	Logger,err := seelog.LoggerFromConfigAsBytes([]byte(config))
	if err != nil {
		log.Fatal(err)
		return
	}
	seelog.ReplaceLogger(Logger)
}

func SwitchLoggerLevel(level string) {
	config := `
<seelog type="asynctimer" asyncinterval="1000000" minlevel="`+level+`" maxlevel="error">
	<outputs formatid="kube">
		<console/>
		<rollingfile formatid="kube" type="size" filename="/var/log/mixer.log" maxsize="1000000" maxrolls="5" />
	</outputs>
	<formats>
		<format id="kube" format="[%Date %Time %LEVEL %FullPath %Func %Line]: %Msg%n"/>
	</formats>
</seelog>`

	Logger,err := seelog.LoggerFromConfigAsBytes([]byte(config))
	if err != nil {
		log.Fatal(err)
		return
	}

	seelog.ReplaceLogger(Logger)
}



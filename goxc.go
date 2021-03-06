package main

import (
	"bufio"
	"strings"

	// "context"

	"io"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"time"

	"errors"
	"fmt"

	"github.com/topxeq/qlang"
	_ "github.com/topxeq/qlang/lib/builtin" // 导入 builtin 包
	_ "github.com/topxeq/qlang/lib/chan"
	"github.com/topxeq/sqltk"

	execq "github.com/topxeq/qlang/exec"
	"github.com/topxeq/qlang/spec"

	// import standard packages
	qlarchivezip "github.com/topxeq/qlang/lib/archive/zip"
	qlbufio "github.com/topxeq/qlang/lib/bufio"
	qlbytes "github.com/topxeq/qlang/lib/bytes"

	qlcrypto "github.com/topxeq/qlang/lib/crypto"
	qlcryptoaes "github.com/topxeq/qlang/lib/crypto/aes"
	qlcryptocipher "github.com/topxeq/qlang/lib/crypto/cipher"
	qlcryptohmac "github.com/topxeq/qlang/lib/crypto/hmac"
	qlcryptomd5 "github.com/topxeq/qlang/lib/crypto/md5"
	qlcryptorand "github.com/topxeq/qlang/lib/crypto/rand"
	qlcryptorsa "github.com/topxeq/qlang/lib/crypto/rsa"
	qlcryptosha1 "github.com/topxeq/qlang/lib/crypto/sha1"
	qlcryptosha256 "github.com/topxeq/qlang/lib/crypto/sha256"
	qlcryptox509 "github.com/topxeq/qlang/lib/crypto/x509"

	qldatabasesql "github.com/topxeq/qlang/lib/database/sql"

	qlencodingbase64 "github.com/topxeq/qlang/lib/encoding/base64"
	qlencodingbinary "github.com/topxeq/qlang/lib/encoding/binary"
	qlencodingcsv "github.com/topxeq/qlang/lib/encoding/csv"
	qlencodinggob "github.com/topxeq/qlang/lib/encoding/gob"
	qlencodinghex "github.com/topxeq/qlang/lib/encoding/hex"
	qlencodingjson "github.com/topxeq/qlang/lib/encoding/json"
	qlencodingpem "github.com/topxeq/qlang/lib/encoding/pem"
	qlencodingxml "github.com/topxeq/qlang/lib/encoding/xml"

	qlerrors "github.com/topxeq/qlang/lib/errors"
	qlflag "github.com/topxeq/qlang/lib/flag"
	qlfmt "github.com/topxeq/qlang/lib/fmt"

	qlhashfnv "github.com/topxeq/qlang/lib/hash/fnv"

	qlhtml "github.com/topxeq/qlang/lib/html"
	qlhtmltemplate "github.com/topxeq/qlang/lib/html/template"

	qlimage "github.com/topxeq/qlang/lib/image"
	qlimage_color "github.com/topxeq/qlang/lib/image/color"
	qlimage_color_palette "github.com/topxeq/qlang/lib/image/color/palette"
	qlimage_draw "github.com/topxeq/qlang/lib/image/draw"
	qlimage_gif "github.com/topxeq/qlang/lib/image/gif"
	qlimage_jpeg "github.com/topxeq/qlang/lib/image/jpeg"
	qlimage_png "github.com/topxeq/qlang/lib/image/png"

	qlio "github.com/topxeq/qlang/lib/io"
	qlio_fs "github.com/topxeq/qlang/lib/io/fs"
	qlioioutil "github.com/topxeq/qlang/lib/io/ioutil"

	qllog "github.com/topxeq/qlang/lib/log"

	qlmath "github.com/topxeq/qlang/lib/math"
	qlmathbig "github.com/topxeq/qlang/lib/math/big"
	qlmathbits "github.com/topxeq/qlang/lib/math/bits"
	qlmathrand "github.com/topxeq/qlang/lib/math/rand"

	qlnet "github.com/topxeq/qlang/lib/net"
	qlnethttp "github.com/topxeq/qlang/lib/net/http"
	qlnet_http_cookiejar "github.com/topxeq/qlang/lib/net/http/cookiejar"
	qlnet_http_httputil "github.com/topxeq/qlang/lib/net/http/httputil"
	qlnet_mail "github.com/topxeq/qlang/lib/net/mail"
	qlnet_rpc "github.com/topxeq/qlang/lib/net/rpc"
	qlnet_rpc_jsonrpc "github.com/topxeq/qlang/lib/net/rpc/jsonrpc"
	qlnet_smtp "github.com/topxeq/qlang/lib/net/smtp"
	qlneturl "github.com/topxeq/qlang/lib/net/url"

	qlos "github.com/topxeq/qlang/lib/os"
	qlos_exec "github.com/topxeq/qlang/lib/os/exec"
	qlos_signal "github.com/topxeq/qlang/lib/os/signal"
	qlos_user "github.com/topxeq/qlang/lib/os/user"

	qlpath "github.com/topxeq/qlang/lib/path"
	qlpathfilepath "github.com/topxeq/qlang/lib/path/filepath"

	qlreflect "github.com/topxeq/qlang/lib/reflect"
	qlregexp "github.com/topxeq/qlang/lib/regexp"
	qlruntime "github.com/topxeq/qlang/lib/runtime"
	qlruntimedebug "github.com/topxeq/qlang/lib/runtime/debug"

	qlsort "github.com/topxeq/qlang/lib/sort"
	qlstrconv "github.com/topxeq/qlang/lib/strconv"
	qlstrings "github.com/topxeq/qlang/lib/strings"
	qlsync "github.com/topxeq/qlang/lib/sync"

	qltext_template "github.com/topxeq/qlang/lib/text/template"
	qltime "github.com/topxeq/qlang/lib/time"

	qlunicode "github.com/topxeq/qlang/lib/unicode"
	qlunicode_utf8 "github.com/topxeq/qlang/lib/unicode/utf8"

	// import 3rd party packages
	qlgithubbeeviketree "github.com/topxeq/qlang/lib/github.com/beevik/etree"
	qlgithubtopxeqimagetk "github.com/topxeq/qlang/lib/github.com/topxeq/imagetk"
	qlgithubtopxeqsqltk "github.com/topxeq/qlang/lib/github.com/topxeq/sqltk"
	qlgithubtopxeqtk "github.com/topxeq/qlang/lib/github.com/topxeq/tk"

	qlgithub_fogleman_gg "github.com/topxeq/qlang/lib/github.com/fogleman/gg"

	qlgithub_360EntSecGroupSkylar_excelize "github.com/topxeq/qlang/lib/github.com/360EntSecGroup-Skylar/excelize"

	qlgithub_kbinani_screenshot "github.com/topxeq/qlang/lib/github.com/kbinani/screenshot"

	qlgithub_stretchr_objx "github.com/topxeq/qlang/lib/github.com/stretchr/objx"

	qlgithub_topxeq_doc2vec_doc2vec "github.com/topxeq/qlang/lib/github.com/topxeq/doc2vec/doc2vec"

	qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi "github.com/topxeq/qlang/lib/github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	// qlgithub_avfs_avfs_fs_memfs "github.com/topxeq/qlang/lib/github.com/avfs/avfs/fs/memfs"
	qlgithub_topxeq_afero "github.com/topxeq/qlang/lib/github.com/topxeq/afero"

	qlgithub_topxeq_socks "github.com/topxeq/qlang/lib/github.com/topxeq/socks"

	qlgithub_topxeq_regexpx "github.com/topxeq/qlang/lib/github.com/topxeq/regexpx"

	qlgithub_topxeq_xmlx "github.com/topxeq/qlang/lib/github.com/topxeq/xmlx"

	qlgithub_topxeq_awsapi "github.com/topxeq/qlang/lib/github.com/topxeq/awsapi"

	qlgithub_cavaliercoder_grab "github.com/topxeq/qlang/lib/github.com/cavaliercoder/grab"

	qlgithub_pterm_pterm "github.com/topxeq/qlang/lib/github.com/pterm/pterm"

	qlgithub_domodwyer_mailyak "github.com/topxeq/qlang/lib/github.com/domodwyer/mailyak"

	

	// full version related start
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"

	// full version related end

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	

	"github.com/topxeq/tk"
)

// Non GUI related

var versionG = "1.80a"

// add tk.ToJSONX

var verboseG = false

var variableG = make(map[string]interface{})

var codeTextG = ""

var qlVMG *qlang.Qlang = nil

var varMutexG sync.Mutex

func exit(argsA ...int) {
	defer func() {
		if r := recover(); r != nil {
			tk.Printfln("exception: %v", r)

			return
		}
	}()

	if argsA == nil || len(argsA) < 1 {
		os.Exit(1)
	}

	os.Exit(argsA[0])
}

func qlEval(strA string) string {
	vmT := qlang.New()

	retG = notFoundG

	errT := vmT.SafeEval(strA)

	if errT != nil {
		return errT.Error()
	}

	rs, ok := vmT.GetVar("outG")

	if ok {
		return tk.Spr("%v", rs)
	}

	if retG != notFoundG {
		return tk.Spr("%v", retG)
	}

	return tk.ErrStrF("no result")
}

func panicIt(valueA interface{}) {
	panic(valueA)
}

func getUint64Value(v reflect.Value) uint16 {
	tk.Pl("%x", v.Interface())

	var p *uint16

	p = (v.Interface().(*uint16))

	return *p
}

func runScript(codeA string, modeA string, argsA ...string) interface{} {

	if modeA == "" || modeA == "0" || modeA == "ql" {
		vmT := qlang.New()

		// if argsA != nil && len(argsA) > 0 {
		vmT.SetVar("argsG", argsA)
		// }

		retG = notFoundG

		errT := vmT.SafeEval(codeA)

		if errT != nil {
			return errT
		}

		rs, ok := vmT.GetVar("outG")

		if ok {
			if rs != nil {
				return rs
			}
		}

		return retG
	} else {
		return tk.SystemCmd("gox", append([]string{codeA}, argsA...)...)
	}

}

func runScriptX(codeA string, argsA ...string) interface{} {

	initQLVM()

	// if argsA != nil && len(argsA) > 0 {
	qlVMG.SetVar("argsG", argsA)
	// }

	retG = notFoundG

	errT := qlVMG.SafeEval(codeA)

	if errT != nil {
		return errT
	}

	rs, ok := qlVMG.GetVar("outG")

	if ok {
		if rs != nil {
			return rs
		}
	}

	return retG

}

func runCode(codeA string, argsA ...string) interface{} {
	initQLVM()

	vmT := qlang.New()

	// if argsA != nil && len(argsA) > 0 {
	vmT.SetVar("argsG", argsA)
	// } else {
	// 	vmT.SetVar("argsG", os.Args)
	// }

	retG = notFoundG

	errT := vmT.SafeEval(codeA)

	if errT != nil {
		return errT
	}

	rs, ok := vmT.GetVar("outG")

	if ok {
		if rs != nil {
			return rs
		}
	}

	if retG != notFoundG {
		return retG
	}

	return retG
}

func getMagic(numberA int) string {
	if numberA < 0 {
		return tk.GenerateErrorString("invalid magic number")
	}

	typeT := numberA % 10

	var fcT string

	if typeT == 8 {
		fcT = tk.DownloadPageUTF8(tk.Spr("https://gitee.com/topxeq/gox/raw/master/magic/%v.gox", numberA), nil, "", 30)

	} else if typeT == 7 {
		fcT = tk.DownloadPageUTF8(tk.Spr("https: //raw.githubusercontent.com/topxeq/gox/master/magic/%v.gox", numberA), nil, "", 30)
	} else {
		return tk.GenerateErrorString("invalid magic number")
	}

	return fcT

}

// native functions 内置函数

var leBufG []string

func leClear() {
	leBufG = make([]string, 0, 100)
}

func leLoadString(strA string) {
	if leBufG == nil {
		leClear()
	}

	leBufG = tk.SplitLines(strA)
}

func leSaveString() string {
	if leBufG == nil {
		leClear()
	}

	return tk.JoinLines(leBufG)
}

func leLoadFile(fileNameA string) error {
	if leBufG == nil {
		leClear()
	}

	strT, errT := tk.LoadStringFromFileE(fileNameA)

	if errT != nil {
		return errT
	}

	leBufG = tk.SplitLines(strT)
	// leBufG, errT = tk.LoadStringListBuffered(fileNameA, false, false)

	return nil
}

func leSaveFile(fileNameA string) error {
	if leBufG == nil {
		leClear()
	}

	var errT error

	textT := tk.JoinLines(leBufG)

	if tk.IsErrStr(textT) {
		return tk.Errf(tk.GetErrStr(textT))
	}

	errT = tk.SaveStringToFileE(textT, fileNameA)

	return errT
}

func leLoadClip() error {
	if leBufG == nil {
		leClear()
	}

	textT := tk.GetClipText()

	if tk.IsErrStr(textT) {
		return tk.Errf(tk.GetErrStr(textT))
	}

	leBufG = tk.SplitLines(textT)

	return nil
}

func leSaveClip() error {
	if leBufG == nil {
		leClear()
	}

	textT := tk.JoinLines(leBufG)

	if tk.IsErrStr(textT) {
		return tk.Errf(tk.GetErrStr(textT))
	}

	return tk.SetClipText(textT)
}

func leViewAll(argsA ...string) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if tk.IfSwitchExistsWhole(argsA, "-nl") {
		textT := tk.JoinLines(leBufG)

		tk.Pln(textT)

	} else {
		for i, v := range leBufG {
			tk.Pl("%v: %v", i, v)
		}
	}

	return nil
}

func leViewLine(idxA int) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if idxA < 0 || idxA >= len(leBufG) {
		return tk.Errf("line index out of range")
	}

	tk.Pln(leBufG[idxA])

	return nil
}

func leGetLine(idxA int) string {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.ErrStrf("buffer not initalized")
	}

	if idxA < 0 || idxA >= len(leBufG) {
		return tk.ErrStrf("line index out of range")
	}

	return leBufG[idxA]
}

func leSetLine(idxA int, strA string) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if idxA < 0 || idxA >= len(leBufG) {
		return tk.Errf("line index out of range")
	}

	leBufG[idxA] = strA

	return nil
}

func leSetLines(startA int, endA int, strA string) error {
	if leBufG == nil {
		leClear()
	}

	if startA > endA {
		return tk.Errf("start index greater than end index")
	}

	listT := tk.SplitLines(strA)

	if endA < 0 {
		rs := make([]string, 0, len(leBufG)+len(listT))

		rs = append(rs, listT...)
		rs = append(rs, leBufG...)

		leBufG = rs

		return nil
	}

	if startA >= len(leBufG) {
		leBufG = append(leBufG, listT...)

		return nil
	}

	if startA < 0 {
		startA = 0
	}

	if endA >= len(leBufG) {
		endA = len(leBufG) - 1
	}

	rs := make([]string, 0, len(leBufG)+len(listT)-1)

	rs = append(rs, leBufG[:startA]...)
	rs = append(rs, listT...)
	rs = append(rs, leBufG[endA+1:]...)

	leBufG = rs

	return nil
}

func leInsertLine(idxA int, strA string) error {
	if leBufG == nil {
		leClear()
	}

	// if leBufG == nil {
	// 	return tk.Errf("buffer not initalized")
	// }

	// if idxA < 0 || idxA >= len(leBufG) {
	// 	return tk.Errf("line index out of range")
	// }

	if idxA < 0 {
		idxA = 0
	}

	listT := tk.SplitLines(strA)

	if idxA >= len(leBufG) {
		leBufG = append(leBufG, listT...)
	} else {
		rs := make([]string, 0, len(leBufG)+1)

		rs = append(rs, leBufG[:idxA]...)
		rs = append(rs, listT...)
		rs = append(rs, leBufG[idxA:]...)

		leBufG = rs

	}

	return nil
}

func leAppendLine(strA string) error {
	if leBufG == nil {
		leClear()
	}

	// if leBufG == nil {
	// 	return tk.Errf("buffer not initalized")
	// }

	// if idxA < 0 || idxA >= len(leBufG) {
	// 	return tk.Errf("line index out of range")
	// }

	listT := tk.SplitLines(strA)

	leBufG = append(leBufG, listT...)

	return nil
}

func leRemoveLine(idxA int) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if idxA < 0 || idxA >= len(leBufG) {
		return tk.Errf("line index out of range")
	}

	rs := make([]string, 0, len(leBufG)+1)

	rs = append(rs, leBufG[:idxA]...)
	rs = append(rs, leBufG[idxA+1:]...)

	leBufG = rs

	return nil
}

func leRemoveLines(startA int, endA int) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if startA < 0 || startA >= len(leBufG) {
		return tk.Errf("start line index out of range")
	}

	if endA < 0 || endA >= len(leBufG) {
		return tk.Errf("end line index out of range")
	}

	if startA > endA {
		return tk.Errf("start line index greater than end line index")
	}

	rs := make([]string, 0, len(leBufG)+1)

	rs = append(rs, leBufG[:startA]...)
	rs = append(rs, leBufG[endA+1:]...)

	leBufG = rs

	return nil
}

func magic(numberA int, argsA ...string) interface{} {
	fcT := getMagic(numberA)

	if tk.IsErrorString(fcT) {
		return tk.ErrorStringToError(fcT)
	}

	return runCode(fcT, argsA...)

}

func NewFuncIntString(funcA *interface{}) *(func(int) string) {
	funcT := (*funcA).(*execq.Function)
	f := func(n int) string {
		return funcT.Call(execq.NewStack(), n).(string)
	}

	return &f
}

func NewFuncFloatString(funcA *interface{}) *(func(float64) string) {
	funcT := (*funcA).(*execq.Function)
	f := func(n float64) string {
		return funcT.Call(execq.NewStack(), n).(string)
	}

	return &f
}

func NewFuncStringString(funcA *interface{}) *(func(string) string) {
	funcT := (*funcA).(*execq.Function)
	f := func(s string) string {
		return funcT.Call(execq.NewStack(), s).(string)
	}

	return &f
}

func NewFuncStringStringB(funcA interface{}) func(string) string {
	funcT := (funcA).(*execq.Function)
	f := func(s string) string {
		return funcT.Call(execq.NewStack(), s).(string)
	}

	return f
}

func NewFuncIntError(funcA *interface{}) *(func(int) error) {
	funcT := (*funcA).(*execq.Function)
	f := func(n int) error {
		return funcT.Call(execq.NewStack(), n).(error)
	}

	return &f
}

func NewFunInterfaceError(funcA *interface{}) *(func(interface{}) error) {
	funcT := (*funcA).(*execq.Function)
	f := func(n interface{}) error {
		return funcT.Call(execq.NewStack(), n).(error)
	}

	return &f
}

func NewFuncStringError(funcA *interface{}) *(func(string) error) {
	funcT := (*funcA).(*execq.Function)
	f := func(s string) error {
		return funcT.Call(execq.NewStack(), s).(error)
	}

	return &f
}

func NewFuncStringStringErrorB(funcA interface{}) func(string) (string, error) {
	funcT := (funcA).(*execq.Function)
	f := func(s string) (string, error) {
		r := funcT.Call(execq.NewStack(), s).([]interface{})

		if r == nil {
			return "", tk.Errf("nil result")
		}

		if len(r) < 2 {
			return "", tk.Errf("incorrect return argument count")
		}

		if r[1] == nil {
			return r[0].(string), nil
		}

		return r[0].(string), r[1].(error)
	}

	return f
}

func NewFuncStringStringError(funcA *interface{}) *(func(string) (string, error)) {
	funcT := (*funcA).(*execq.Function)
	f := func(s string) (string, error) {
		r := funcT.Call(execq.NewStack(), s).([]interface{})

		if r == nil {
			return "", tk.Errf("nil result")
		}

		if len(r) < 2 {
			return "", tk.Errf("incorrect return argument count")
		}

		if r[1] == nil {
			return r[0].(string), nil
		}

		return r[0].(string), r[1].(error)
	}

	return &f
}

func NewFuncInterfaceInterfaceErrorB(funcA interface{}) func(interface{}) (interface{}, error) {
	funcT := (funcA).(*execq.Function)
	f := func(s interface{}) (interface{}, error) {
		r := funcT.Call(execq.NewStack(), s).([]interface{})

		if r[1] == nil {
			return r[0].(interface{}), nil
		}

		return r[0].(interface{}), r[1].(error)
	}

	return f
}

func NewFuncInterfaceInterfaceError(funcA *interface{}) *(func(interface{}) (interface{}, error)) {
	funcT := (*funcA).(*execq.Function)
	f := func(s interface{}) (interface{}, error) {
		r := funcT.Call(execq.NewStack(), s).([]interface{})

		// if r == nil {
		// 	return "", tk.Errf("nil result")
		// }

		// if len(r) < 2 {
		// 	return "", tk.Errf("incorrect return argument count")
		// }

		if r[1] == nil {
			return r[0].(interface{}), nil
		}

		return r[0].(interface{}), r[1].(error)
	}

	return &f
}

func NewFuncInterfaceError(funcA *interface{}) *(func(interface{}) error) {
	funcT := (*funcA).(*execq.Function)
	f := func(s interface{}) error {
		return funcT.Call(execq.NewStack(), s).(error)
	}

	return &f
}

func NewFunc(funcA *interface{}) *(func()) {
	funcT := (*funcA).(*execq.Function)
	f := func() {
		funcT.Call(execq.NewStack())

		return
	}

	return &f
}

func NewFuncB(funcA interface{}) func() {
	funcT := (funcA).(*execq.Function)
	f := func() {
		funcT.Call(execq.NewStack())

		return
	}

	return f
}

func NewFuncError(funcA *interface{}) *(func() error) {
	funcT := (*funcA).(*execq.Function)
	f := func() error {
		return funcT.Call(execq.NewStack()).(error)
	}

	return &f
}

func NewFuncInterface(funcA *interface{}) *(func() interface{}) {
	funcT := (*funcA).(*execq.Function)
	f := func() interface{} {
		return funcT.Call(execq.NewStack()).(interface{})
	}

	return &f
}

func NewFuncIntStringError(funcA *interface{}) *(func(int) (string, error)) {
	funcT := (*funcA).(*execq.Function)
	f := func(n int) (string, error) {
		r := funcT.Call(execq.NewStack(), n).([]interface{})

		if r == nil {
			return "", tk.Errf("nil result")
		}

		if len(r) < 2 {
			return "", tk.Errf("incorrect return argument count")
		}

		if r[1] == nil {
			return r[0].(string), nil
		}

		return r[0].(string), r[1].(error)
	}

	return &f
}

func NewFuncFloatStringError(funcA *interface{}) *(func(float64) (string, error)) {
	funcT := (*funcA).(*execq.Function)
	f := func(n float64) (string, error) {
		r := funcT.Call(execq.NewStack(), n).([]interface{})

		if r == nil {
			return "", tk.Errf("nil result")
		}

		if len(r) < 2 {
			return "", tk.Errf("incorrect return argument count")
		}

		if r[1] == nil {
			return r[0].(string), nil
		}

		return r[0].(string), r[1].(error)
	}

	return &f
}

func printValue(nameA string) {

	v, ok := qlVMG.GetVar(nameA)

	if !ok {
		tk.Pl("no variable by the name found: %v", nameA)
		return
	}

	tk.Pl("%v(%T): %v", nameA, v, v)

}

func defined(nameA string) bool {

	_, ok := qlVMG.GetVar(nameA)

	return ok

}

func nilToEmpty(vA interface{}, argsA ...string) string {

	if vA == nil {
		return ""
	}

	if vA == spec.Undefined {
		return ""
	}

	if tk.IsNil(vA) {
		return ""
	}

	if (argsA != nil) && (len(argsA) > 0) {
		vf, ok := vA.(float64)
		if ok {
			if tk.IfSwitchExistsWhole(argsA, "-nofloat") {
				return tk.ToStr(int(vf))
			} else {
				return tk.Float64ToStr(vA.(float64))
			}
		}
	}

	return fmt.Sprintf("%v", vA)

}

func isValid(vA interface{}, argsA ...string) bool {

	if vA == nil {
		return false
	}

	if vA == spec.Undefined {
		return false
	}

	if tk.IsNil(vA) {
		return false
	}

	if (argsA != nil) && (len(argsA) > 0) {
		typeT := fmt.Sprintf("%T", vA)

		if typeT == argsA[0] {
			return true
		} else {
			return false
		}
	}

	return true
}

func isDefined(vA interface{}) bool {
	if vA == spec.Undefined {
		return false
	}

	return true
}

func logPrint(formatA string, argsA ...interface{}) {
	tk.Pl(formatA, argsA...)
	tk.LogWithTimeCompact(formatA, argsA...)
}

// -1 return random item
func getArrayItem(aryA interface{}, idxA int, defaultA ...interface{}) interface{} {
	var hasDefaultT = false
	if len(defaultA) > 0 {
		hasDefaultT = true
	}

	if aryA == nil {
		if hasDefaultT {
			return defaultA[0]
		}

		return ""
	}

	switch aryT := aryA.(type) {
	case []interface{}:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	case []string:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	case []int:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	case []float64:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	case []bool:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	}

	return ""

}

func getMapItem(mapA interface{}, keyA string, defaultA ...interface{}) interface{} {
	var hasDefaultT = false
	if len(defaultA) > 0 {
		hasDefaultT = true
	}

	if mapA == nil {
		if hasDefaultT {
			return defaultA[0]
		}

		return ""
	}

	switch mapT := mapA.(type) {
	case map[string]interface{}:
		itemT, ok := mapT[keyA]
		if !ok {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		return itemT
	case map[string]string:
		itemT, ok := mapT[keyA]
		if !ok {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		return itemT
	}

	return ""
}

func strJoin(aryA interface{}, sepA string, defaultA ...string) string {
	var defaultT string = ""
	if (defaultA != nil) && (len(defaultA) > 0) {
		defaultT = defaultA[0]
	}

	if aryA == nil {
		return defaultT
	}

	switch v := aryA.(type) {
	case []string:
		return strings.Join(v, sepA)
	case []interface{}:
		var bufT strings.Builder
		for j, jv := range v {
			if j > 0 {
				bufT.WriteString(sepA)
			}

			bufT.WriteString(fmt.Sprintf("%v", jv))
		}

		return bufT.String()
	}

	return defaultT
}

func strToTime(strA string, formatA ...string) interface{} {
	formatT := tk.TimeFormat

	if (formatA != nil) && (len(formatA) > 0) {
		formatT = formatA[0]
	}

	timeT, errT := tk.StrToTimeByFormat(strA, formatT)

	if errT != nil {
		return spec.Undefined
	}

	return timeT
}



var scriptPathG string

func importQLNonGUIPackages() {
	// getPointer := func(nameA string) {

	// 	v, ok := qlVMG.GetVar(nameA)

	// 	if !ok {
	// 		tk.Pl("no variable by the name found: %v", nameA)
	// 		return
	// 	}

	// 	tk.Pl("%v(%T): %v", nameA, v, v)

	// }

	// setString := func(p *string, strA string) {
	// 	*p = strA
	// }

	// import native functions and global variables 内置函数与全局变量
	var defaultExports = map[string]interface{}{
		// 其中 tk.开头的函数都是github.com/topxeq/tk包中的，可以去pkg.go.dev/github.com/topxeq/tk查看函数定义

		// common related 一般函数
		"defined":       defined,               // 查看某变量是否已经定义，注意参数是字符串类型的变量名，例： if defined("a") {...}
		"pass":          tk.Pass,               // 没有任何操作的函数，一般用于脚本结尾避免脚本返回一个结果导致输出乱了
		"isDefined":     isDefined,             // 判断某变量是否已经定义，与defined的区别是传递的是变量名而不是字符串方式的变量，例： if isDefined(a) {...}
		"isValid":       isValid,               // 判断某变量是否已经定义，并且不是nil或空字符串，如果传入了第二个参数，还可以判断该变量是否类型是该类型，例： if isValid(a, "string") {...}
		"eval":          qlEval,                // 运行一段Gox语言代码
		"typeOf":        tk.TypeOfValue,        // 给出某变量的类型名
		"typeOfReflect": tk.TypeOfValueReflect, // 给出某变量的类型名（使用了反射方式）
		"exit":          tk.Exit,               // 立即退出脚本的执行，可以带一个整数作为参数，也可以没有
		"setValue":      tk.SetValue,           // 用反射的方式设定一个变量的值
		"getValue":      tk.GetValue,           // 用反射的方式获取一个变量的值
		"setVar":        tk.SetVar,             // 设置一个全局变量，例： setVar("a", "value of a")
		"getVar":        tk.GetVar,             // 获取一个全局变量的值，例： v = getVar("a")
		"isNil":         tk.IsNil,              // 判断一个变量或表达式是否为nil
		"deepClone":     tk.DeepClone,
		"deepCopy":      tk.DeepCopyFromTo,
		"run":           runFile,
		"runCode":       runCode,
		"runScript":     runScript,
		"magic":         magic,

		// output related 输出相关
		"pv":        printValue,   // 输出一个变量的值，注意参数是字符串类型的变量名，例： pv("a")
		"pr":        tk.Pr,        // 等同于其他语言中的print
		"prf":       tk.Printf,    // 等同于其他语言中的printf
		"pln":       tk.Pln,       // 等同于其他语言中的println
		"printfln":  tk.Pl,        // 等同于其他语言中的printf，但多输出一个回车换行
		"pl":        tk.Pl,        // 等同于printfln
		"sprintf":   fmt.Sprintf,  // 等同于其他语言中的sprintf
		"spr":       fmt.Sprintf,  // 等同于sprintf
		"fprintf":   fmt.Fprintf,  // 等同于其他语言中的frintf
		"plv":       tk.Plv,       // 输出某变量或表达式的内容/值，以Go语言内部的表达方式，例如字符串将加上双引号
		"plvx":      tk.Plvx,      // 输出某变量或表达式的内容/值和类型等信息
		"plNow":     tk.PlNow,     // 相当于pl，但前面多加了一个时间标记
		"plVerbose": tk.PlVerbose, // 相当于pl，但前面多了一个布尔类型的参数，可以传入一个verbose变量，指定是否输出该信息，例：
		// v = false
		// plVerbose(v, "a: %v", 3) // 由于v的值为false，因此本条语句将不输出
		"plvsr":  tk.Plvsr,     // 输出多个变量或表达式的值，每行一个
		"plerr":  tk.PlErr,     // 快捷输出一个error类型的值
		"plExit": tk.PlAndExit, // 相当于pl然后exit退出脚本的执行

		// input related 输入相关
		"getInput":     tk.GetUserInput,      // 从命令行获取用户的输入
		"getInputf":    tk.GetInputf,         // 从命令行获取用户的输入，同时可以用printf先输出一个提示信息
		"getPasswordf": tk.GetInputPasswordf, // 从命令行获取密码输入，输入信息将不显示

		// math related数学相关
		"bitXor": tk.BitXor, // 异或运算

		// string related 字符串相关
		"trim":             tk.Trim,                   // 取出字符串前后的空白字符
		"strTrim":          tk.Trim,                   // 等同于trim
		"toLower":          strings.ToLower,           // 字符串转小写
		"toUpper":          strings.ToUpper,           // 字符串转大写
		"strContains":      strings.Contains,          // 判断字符串中是否包含某个字串
		"strReplace":       tk.Replace,                // 替换字符串中的字串
		"strJoin":          strJoin,                   // 连接一个字符串数组，以指定的分隔符，例： s = strJoin(listT, "\n")
		"strSplit":         strings.Split,             // 拆分一个字符串为数组，例： listT = strSplit(strT, "\n")
		"splitLines":       tk.SplitLines,             // 相当于strSplit(strT, "\n")
		"startsWith":       tk.StartsWith,             // 判断字符串是否以某子串开头
		"strStartsWith":    tk.StartsWith,             // 等同于startsWith
		"endsWith":         tk.EndsWith,               // 判断字符串是否以某子串结尾
		"strEndsWith":      tk.EndsWith,               // 等同于endsWith
		"strIn":            tk.InStrings,              // 判断字符串是否在一个字符串列表中出现，函数定义： strIn(strA string, argsA ...string) bool
		"getNowStr":        tk.GetNowTimeStringFormal, // 获取一个表示当前时间的字符串，格式：2020-02-02 08:09:15
		"getNowStrCompact": tk.GetNowTimeString,       // 获取一个简化的表示当前时间的字符串，格式：20200202080915
		"genRandomStr":     tk.GenerateRandomString,   // 生成随机字符串，函数定义： (minCharA, maxCharA int, hasUpperA, hasLowerA, hasDigitA, hasSpecialCharA, hasSpaceA bool, hasInvalidChars bool) string

		// regex related 正则表达式相关
		"regMatch":     tk.RegMatchX,          // 判断某字符串是否完整符合某表达式，例： if regMatch(mailT, `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`) {...}
		"regContains":  tk.RegContainsX,       // 判断某字符串是否包含符合正则表达式的子串，例： if regContains("abccd", "b.c") {...}
		"regFind":      tk.RegFindFirstX,      // 根据正则表达式在字符串中寻找第一个匹配，函数定义： func regFind(strA, patternA string, groupA int) string
		"regFindAll":   tk.RegFindAllX,        // 根据正则表达式在字符串中寻找所有匹配，函数定义： func regFindAll(strA, patternA string, groupA int) []string
		"regFindIndex": tk.RegFindFirstIndexX, // 根据正则表达式在字符串中第一个匹配的为止，函数定义： func regFindIndex(strA, patternA string) (int, int)
		"regReplace":   tk.RegReplaceX,        // 根据正则表达式在字符串中进行替换，函数定义： regReplace(strA, patternA, replaceA string) string
		"regSplit":     tk.RegSplitX,          // 根据正则表达式分割字符串（以符合条件的匹配来分割），函数定义： regSplit(strA, patternA string, nA ...int) []string

		// conversion related 转换相关
		"nilToEmpty": nilToEmpty,                  // 将nil等值都转换为空字符串
		"intToStr":   tk.IntToStrX,                // 整数转字符串
		"strToInt":   tk.StrToIntWithDefaultValue, // 字符串转整数
		"floatToStr": tk.Float64ToStr,             // 浮点数转字符串
		"timeToStr":  tk.FormatTime,               // 时间转字符串，函数定义: timeToStr(timeA time.Time, formatA ...string) string
		"strToTime":  strToTime,                   // 字符串转时间
		"toStr":      tk.ToStr,                    // 任意值转字符串
		"toInt":      tk.ToInt,                    // 任意值转整数
		"toFloat":    tk.ToFloat,                  // 任意值转浮点数

		// array/map related 数组（切片）/映射（字典）相关
		"remove":       tk.RemoveItemsInArray,               // 从切片中删除指定的项，例： remove(aryT, 3)
		"getMapString": tk.SafelyGetStringForKeyWithDefault, // 从映射中获得指定的键值，避免返回nil，函数定义：func getMapString(mapA map[string]string, keyA string, defaultA ...string) string， 不指定defaultA将返回空字符串
		"getMapItem":   getMapItem,                          // 类似于getMapString，但可以取任意类型的值
		"getArrayItem": getArrayItem,                        // 类似于getMapItem，但是是去一个切片中指定序号的值
		"joinList":     tk.JoinList,                         // 类似于strJoin，但可以连接任意类型的值

		// error related 错误处理相关
		"isError":          tk.IsError,          // 判断表达式的值是否为error类型
		"isErrStr":         tk.IsErrStr,         // 判断字符串是否是TXERROR:开始的字符串
		"checkError":       tk.CheckError,       // 检查变量，如果是error则立即停止脚本的执行
		"checkErrorString": tk.CheckErrorString, // 检查变量，如果是TXERROR:开始的字符串则立即停止脚本的执行
		"checkErrStr":      tk.CheckErrStr,      // 等同于checkErrorString
		"checkErrf":        tk.CheckErrf,        // 检查变量，如果是error则立即停止脚本的执行，之前可以printfln输出信息
		"checkErrStrf":     tk.CheckErrStrf,     // 检查变量，如果是TXERROR:开始的字符串则立即停止脚本的执行，之前可以printfln输出信息
		"fatalf":           tk.Fatalf,           // printfln输出信息后终止脚本的执行
		"errStr":           tk.ErrStr,           // 生成TXERROR:开始的字符串
		"errStrf":          tk.ErrStrF,          // 生成TXERROR:开始的字符串，类似sprintf的用法
		"getErrStr":        tk.GetErrStr,        // 从TXERROR:开始的字符串获取其后的错误信息
		"errf":             tk.Errf,             // 生成error类型的变量，其中提示信息类似sprintf的用法

		// encode/decode related 编码/解码相关
		"xmlEncode":          tk.EncodeToXMLString,    // 编码为XML
		"xmlDecode":          tk.FromXMLWithDefault,   // 解码XML为对象，函数定义：(xmlA string, defaultA interface{}) interface{}
		"htmlEncode":         tk.EncodeHTML,           // HTML编码（&nbsp;等）
		"htmlDecode":         tk.DecodeHTML,           // HTML解码
		"base64Encode":       tk.EncodeToBase64,       // Base64编码，输入参数是[]byte字节数组
		"base64Decode":       tk.DecodeFromBase64,     // base64解码
		"md5Encode":          tk.MD5Encrypt,           // MD5编码
		"md5":                tk.MD5Encrypt,           // 等同于md5Encode
		"hexEncode":          tk.StrToHex,             // 16进制编码
		"strToHex":           tk.StrToHex,             // 等同于hexEncode
		"hexDecode":          tk.HexToStr,             // 16进制解码
		"hexToStr":           tk.HexToStr,             // 等同于hexDecode
		"jsonEncode":         tk.ObjectToJSON,         // JSON编码
		"jsonDecode":         tk.JSONToObject,         // JSON解码
		"toJSON":             tk.ToJSONX,              // 增强的JSON编码，建议使用，函数定义： toJSON(objA interface{}, optsA ...string) string，参数optsA可选。例：s = toJSON(textA, "-indent", "-sort")
		"fromJSON":           tk.FromJSONWithDefault,  // 增强的JSON解码，建议使用，函数定义： fromJSON(jsonA string, defaultA ...interface{}) interface{}
		"simpleEncode":       tk.EncodeStringCustomEx, // 简单编码，主要为了文件名和网址名不含非法字符
		"simpleDecode":       tk.DecodeStringCustom,   // 简单编码的解码，主要为了文件名和网址名不含非法字符
		"tableToMSSArray":    tk.TableToMSSArray,
		"tableToMSSMap":      tk.TableToMSSMap,
		"tableToMSSMapArray": tk.TKX.TableToMSSMapArray,

		// encrypt/decrypt related 加密/解密相关
		"encryptStr":  tk.EncryptStringByTXDEF, // 加密字符串，第二个参数（可选）是密钥字串
		"decryptStr":  tk.DecryptStringByTXDEF, // 解密字符串，第二个参数（可选）是密钥字串
		"encryptData": tk.EncryptDataByTXDEF,   // 加密二进制数据（[]byte类型），第二个参数（可选）是密钥字串
		"decryptData": tk.DecryptDataByTXDEF,   // 解密二进制数据（[]byte类型），第二个参数（可选）是密钥字串

		// log related 日志相关
		"setLogFile": tk.SetLogFile,         // 设置日志文件路径，下面有关日志的函数将用到
		"logf":       tk.LogWithTimeCompact, // 输出到日志文件，函数定义： func logf(formatA string, argsA ...interface{})
		"logPrint":   logPrint,              // 同时输出到标准输出和日志文件

		// system related 系统相关
		"getClipText":  tk.GetClipText,        // 从系统剪贴板获取文本，例： textT = getClipText()
		"setClipText":  tk.SetClipText,        // 设定系统剪贴板中的文本，例： setClipText("测试")
		"systemCmd":    tk.SystemCmd,          // 执行一条系统命令，例如： systemCmd("cmd", "/k", "copy a.txt b.txt")
		"ifFileExists": tk.IfFileExists,       // 判断文件是否存在
		"fileExists":   tk.IfFileExists,       // 等同于ifFileExists
		"joinPath":     filepath.Join,         // 连接文件路径，等同于Go语言标准库中的path/filepath.Join
		"getFileSize":  tk.GetFileSizeCompact, // 获取文件大小
		"getFileList":  tk.GetFileList,        // 获取指定目录下的符合条件的所有文件，例：listT = getFileList(pathT, "-recursive", "-pattern=*", "-exclusive=*.txt", "-verbose")
		"loadText":     tk.LoadStringFromFile, // 从文件中读取文本字符串，函数定义：func loadText(fileNameA string) string，出错时返回TXERROR:开头的字符串指明原因
		"saveText":     tk.SaveStringToFile,   // 将字符串保存到文件，函数定义： func saveText(strA string, fileA string) string
		"loadBytes":    tk.LoadBytesFromFileE, // 从文件中读取二进制数据，函数定义：func loadBytes(fileNameA string, numA ...int) ([]byte, error)
		"saveBytes":    tk.SaveBytesToFile,    // 将二进制数据保存到文件，函数定义： func saveBytes(bytesA []byte, fileA string) string
		"sleep":        tk.SleepSeconds,       // 休眠指定的秒数，例：sleep(30)
		"sleepSeconds": tk.SleepSeconds,       // 等同于sleep

		// command-line 命令行处理相关
		"getParameter":   tk.GetParameterByIndexWithDefaultValue, // 按顺序序号获取命令行参数，其中0代表第一个参数，也就是软件名称或者命令名称，1开始才是第一个参数，注意参数不包括开关，即类似-verbose=true这样的，函数定义：func getParameter(argsA []string, idxA int, defaultA string) string
		"getSwitch":      tk.GetSwitchWithDefaultValue,           // 获取命令行参数中的开关，用法：tmps = getSwitch(args, "-verbose=", "false")，第三个参数是默认值（如果在命令行中没取到的话返回该值）
		"getIntSwitch":   tk.GetSwitchWithDefaultIntValue,        // 与getSwitch类似，但获取到的是整数的值
		"switchExists":   tk.IfSwitchExistsWhole,                 // 判断命令行参数中是否存在开关（完整的，），用法：flag = switchExists(args, "-restart")
		"ifSwitchExists": tk.IfSwitchExistsWhole,                 // 等同于switchExists

		// network related 网络相关
		"newSSHClient": tk.NewSSHClient, // 新建一个SSH连接，以便执行各种SSH操作，例：
		// clientT, errT = newSSHClient(hostName, port, userName, password)
		// outT, errT = clientT.Run(`ls -p; cat abc.txt`)
		// errT = clientT.Upload(`./abc.txt`, strReplace(joinPath(pathT, `abc.txt`), `\`, "/"))
		// errT = clientT.Download(`down.txt`, `./down.txt`)
		"mapToPostData": tk.MapToPostData,    // 从一个映射（map）对象生成进行POST请求的参数对象，函数定义func mapToPostData(postDataA map[string]string) url.Values
		"getWebPage":    tk.DownloadPageUTF8, // 进行一个网络HTTP请求并获得服务器返回结果，或者下载一个网页，函数定义func getWebPage(urlA string, postDataA url.Values, customHeaders string, timeoutSecsA time.Duration, optsA ...string) string
		// customHeadersA 是自定义请求头，内容是多行文本形如 charset: utf-8。如果冒号后还有冒号，要替换成`
		// 返回结果是TXERROR字符串，即如果是以TXERROR:开头，则表示错误信息，否则是网页或请求响应
		"downloadFile": tk.DownloadFile, // 从网络下载一个文件，函数定义func downloadFile(urlA, dirA, fileNameA string, argsA ...string) string
		"httpRequest":  tk.RequestX,     // 进行一个网络HTTP请求并获得服务器返回结果，函数定义func httpRequest(urlA, methodA, reqBodyA string, customHeadersA string, timeoutSecsA time.Duration, optsA ...string) (string, error)
		// 其中methodA可以是"GET"，"POST"等
		// customHeadersA 是自定义请求头，内容是多行文本形如 charset: utf-8。如果冒号后还有冒号，要替换成`
		"getFormValue":         tk.GetFormValueWithDefaultValue,  // 从HTTP请求中获取字段参数，可以是Query参数，也可以是POST参数，函数定义func getFormValue(reqA *http.Request, keyA string, defaultA string) string
		"formValueExist":       tk.IfFormValueExists,             // 判断HTTP请求中的是否有某个字段参数，函数定义func formValueExist(reqA *http.Request, keyA string) bool
		"ifFormValueExist":     tk.IfFormValueExists,             // 等同于formValueExist
		"formToMap":            tk.FormToMap,                     // 将HTTP请求中的form内容转换为map（字典/映射类型），例：mapT = formToMap(req.Form)
		"generateJSONResponse": tk.GenerateJSONPResponseWithMore, // 生成Web API服务器的JSON响应，支持JSONP，例：return generateJSONResponse("fail", sprintf("数据库操作失败：%v", errT), req)

		// database related
		"dbConnect": sqltk.ConnectDBX, // 连接数据库以便后续读写操作，例：
		// dbT = dbConnect("sqlserver", "server=127.0.0.1;port=1443;portNumber=1443;user id=user;password=userpass;database=db1")
		// 	if isError(dbT) {
		// 		fatalf("打开数据库%v错误：%v", dbT)
		// 	}
		// }
		// defer dbT.Close()

		"dbExec": sqltk.ExecDBX, // 进行数据库操作，例：
		// rs := dbExec(dbT, `insert into table1 (field1,id,field2) values('value1',1,'value2')`
		// 	if isError(rs) {
		// 		fatalf("新增数据库记录时发生错误：%v", rs)
		// 	}
		// }
		// insertID, affectedRows = rs[0], rs[1]
		"dbQuery": sqltk.QueryDBX, // 进行数据库查询，所有字段结果都将转换为字符串，返回结果为[]map[string]string，用JSON格式表达类似：[{"Field1": "Value1", "Field2": "Value2"},{"Field1": "Value1a", "Field2": "Value2a"}]，例：
		// sqlRsT = dbQuery(dbT, `SELECT * FROM TABLE1 WHERE ID=3`)
		// if isError(sqlRsT) {
		//		fatalf("查询数据库错误：%v", dbT)
		//	}
		// pl("在数据库中找到%v条记录", len(sqlRsT))
		"dbQueryCount": sqltk.QueryCountX, // 与dbQuery类似，但主要进行数量查询，也支持结果只有一个整数的查询，例：
		// sqlRsT = dbQueryCount(dbT, `SELECT COUNT(*) FROM TABLE1 WHERE ID>3`)
		// if isError(sqlRsT) {
		//		fatalf("查询数据库错误：%v", dbT)
		//	}
		// pl("在数据库中共有符合条件的%v条记录", sqlRsT)
		"dbQueryString": sqltk.QueryStringX, // 与dbQueryCount类似，但主要支持结果只有一个字符串的查询

		// line editor related 内置行文本编辑器有关
		"leClear":       leClear,       // 清空行文本编辑器缓冲区，例：leClear()
		"leLoadStr":     leLoadString,  // 行文本编辑器缓冲区载入指定字符串内容，例：leLoadStr("abc\nbbb\n结束")
		"leSetAll":      leLoadString,  // 等同于leLoadString
		"leSaveStr":     leSaveString,  // 取出行文本编辑器缓冲区中内容，例：s = leSaveStr()
		"leGetAll":      leSaveString,  // 等同于leSaveStr
		"leLoad":        leLoadFile,    // 从文件中载入文本到行文本编辑器缓冲区中，例：err = leLoad(`c:\test.txt`)
		"leLoadFile":    leLoadFile,    // 等同于leLoad
		"leSave":        leSaveFile,    // 将行文本编辑器缓冲区中内容保存到文件中，例：err = leSave(`c:\test.txt`)
		"leSaveFile":    leSaveFile,    // 等同于leSave
		"leLoadClip":    leLoadClip,    // 从剪贴板中载入文本到行文本编辑器缓冲区中，例：err = leLoadClip()
		"leSaveClip":    leSaveClip,    // 将行文本编辑器缓冲区中内容保存到剪贴板中，例：err = leSaveClip()
		"leInsert":      leInsertLine,  // 行文本编辑器缓冲区中的指定位置前插入指定内容，例：err = leInsert(3， "abc")
		"leInsertLine":  leInsertLine,  // 行文本编辑器缓冲区中的指定位置前插入指定内容，例：err = leInsertLine(3， "abc")
		"leAppend":      leAppendLine,  // 行文本编辑器缓冲区中的指定位置后插入指定内容，例：err = leAppend(3， "abc")
		"leAppendLine":  leAppendLine,  // 行文本编辑器缓冲区中的指定位置后插入指定内容，例：err = leAppendLine(3， "abc")
		"leSet":         leSetLine,     // 设定行文本编辑器缓冲区中的指定行为指定内容，例：err = leSet(3， "abc")
		"leSetLine":     leSetLine,     // 设定行文本编辑器缓冲区中的指定行为指定内容，例：err = leSetLine(3， "abc")
		"leSetLines":    leSetLines,    // 设定行文本编辑器缓冲区中指定范围的多行为指定内容，例：err = leSetLines(3, 5， "abc\nbbb")
		"leRemove":      leRemoveLine,  // 删除行文本编辑器缓冲区中的指定行，例：err = leRemove(3)
		"leRemoveLine":  leRemoveLine,  // 删除行文本编辑器缓冲区中的指定行，例：err = leRemoveLine(3)
		"leRemoveLines": leRemoveLines, // 删除行文本编辑器缓冲区中指定范围的多行，例：err = leRemoveLines(1, 3)
		"leViewAll":     leViewAll,     // 查看行文本编辑器缓冲区中的所有内容，例：allText = leViewAll()
		"leView":        leViewLine,    // 查看行文本编辑器缓冲区中的指定行，例：lineText = leView(18)

		

		// misc 杂项函数
		"newFunc":    NewFuncB,                        // 将Gox语言中的定义的函数转换为Go语言中类似 func f() 的形式
		"newFuncIIE": NewFuncInterfaceInterfaceErrorB, // 将Gox语言中的定义的函数转换为Go语言中类似 func f(a interface{}) (interface{}, error) 的形式
		"newFuncSSE": NewFuncStringStringErrorB,       // 将Gox语言中的定义的函数转换为Go语言中类似 func f(a string) (string, error) 的形式
		"newFuncSS":  NewFuncStringStringB,            // 将Gox语言中的定义的函数转换为Go语言中类似 func f(a string) string 的形式

		// global variables 全局变量
		"timeFormatG":        tk.TimeFormat,        // 用于时间处理时的时间格式，值为"2006-01-02 15:04:05"
		"timeFormatCompactG": tk.TimeFormatCompact, // 用于时间处理时的简化时间格式，值为"20060102150405"

		"scriptPathG": scriptPathG, // 所执行脚本的路径
		"versionG":    versionG,    // Gox/Goxc的版本号
		"leBufG":      leBufG,      // 内置行文本编辑器所用的编辑缓冲区

		
	}

	qlang.Import("", defaultExports)

	var imiscExports = map[string]interface{}{
		"NewFunc":                         NewFunc,
		"NewFuncError":                    NewFuncError,
		"NewFuncInterface":                NewFuncInterface,
		"NewFuncInterfaceError":           NewFuncInterfaceError,
		"NewFuncInterfaceInterfaceError":  NewFuncInterfaceInterfaceError,
		"NewFuncInterfaceInterfaceErrorB": NewFuncInterfaceInterfaceErrorB,
		"NewFuncIntString":                NewFuncIntString,
		"NewFuncIntError":                 NewFuncIntError,
		"NewFuncFloatString":              NewFuncFloatString,
		"NewFuncFloatStringError":         NewFuncFloatStringError,
		"NewFuncStringString":             NewFuncStringString,
		"NewFuncStringError":              NewFuncStringError,
		"NewFuncStringStringError":        NewFuncStringStringError,
		"NewFuncStringStringErrorB":       NewFuncStringStringErrorB,
		"NewFuncIntStringError":           NewFuncIntStringError,
	}

	qlang.Import("imisc", imiscExports)

	qlang.Import("archive_zip", qlarchivezip.Exports)
	qlang.Import("bufio", qlbufio.Exports)
	qlang.Import("bytes", qlbytes.Exports)

	qlang.Import("crypto", qlcrypto.Exports)
	qlang.Import("crypto_aes", qlcryptoaes.Exports)
	qlang.Import("crypto_cipher", qlcryptocipher.Exports)
	qlang.Import("crypto_hmac", qlcryptohmac.Exports)
	qlang.Import("crypto_md5", qlcryptomd5.Exports)
	qlang.Import("crypto_rand", qlcryptorand.Exports)
	qlang.Import("crypto_rsa", qlcryptorsa.Exports)
	qlang.Import("crypto_sha256", qlcryptosha256.Exports)
	qlang.Import("crypto_sha1", qlcryptosha1.Exports)
	qlang.Import("crypto_x509", qlcryptox509.Exports)

	qlang.Import("database_sql", qldatabasesql.Exports)

	qlang.Import("encoding_pem", qlencodingpem.Exports)
	qlang.Import("encoding_base64", qlencodingbase64.Exports)
	qlang.Import("encoding_binary", qlencodingbinary.Exports)
	qlang.Import("encoding_csv", qlencodingcsv.Exports)
	qlang.Import("encoding_gob", qlencodinggob.Exports)
	qlang.Import("encoding_hex", qlencodinghex.Exports)
	qlang.Import("encoding_json", qlencodingjson.Exports)
	qlang.Import("encoding_xml", qlencodingxml.Exports)

	qlang.Import("errors", qlerrors.Exports)

	qlang.Import("flag", qlflag.Exports)
	qlang.Import("fmt", qlfmt.Exports)

	qlang.Import("hash_fnv", qlhashfnv.Exports)

	qlang.Import("html", qlhtml.Exports)
	qlang.Import("html_template", qlhtmltemplate.Exports)

	qlang.Import("image", qlimage.Exports)
	qlang.Import("image_color", qlimage_color.Exports)
	qlang.Import("image_color_palette", qlimage_color_palette.Exports)
	qlang.Import("image_draw", qlimage_draw.Exports)
	qlang.Import("image_gif", qlimage_gif.Exports)
	qlang.Import("image_jpeg", qlimage_jpeg.Exports)
	qlang.Import("image_png", qlimage_png.Exports)

	qlang.Import("io", qlio.Exports)
	qlang.Import("io_ioutil", qlioioutil.Exports)
	qlang.Import("io_fs", qlio_fs.Exports)

	qlang.Import("log", qllog.Exports)

	qlang.Import("math", qlmath.Exports)
	qlang.Import("math_big", qlmathbig.Exports)
	qlang.Import("math_bits", qlmathbits.Exports)
	qlang.Import("math_rand", qlmathrand.Exports)

	qlang.Import("net", qlnet.Exports)
	qlang.Import("net_http", qlnethttp.Exports)
	qlang.Import("net_http_cookiejar", qlnet_http_cookiejar.Exports)
	qlang.Import("net_http_httputil", qlnet_http_httputil.Exports)
	qlang.Import("net_mail", qlnet_mail.Exports)
	qlang.Import("net_rpc", qlnet_rpc.Exports)
	qlang.Import("net_rpc_jsonrpc", qlnet_rpc_jsonrpc.Exports)
	qlang.Import("net_smtp", qlnet_smtp.Exports)
	qlang.Import("net_url", qlneturl.Exports)

	qlang.Import("os", qlos.Exports)
	qlang.Import("os_exec", qlos_exec.Exports)
	qlang.Import("os_signal", qlos_signal.Exports)
	qlang.Import("os_user", qlos_user.Exports)
	qlang.Import("path", qlpath.Exports)
	qlang.Import("path_filepath", qlpathfilepath.Exports)

	qlang.Import("reflect", qlreflect.Exports)
	qlang.Import("regexp", qlregexp.Exports)

	qlang.Import("runtime", qlruntime.Exports)
	qlang.Import("runtime_debug", qlruntimedebug.Exports)

	qlang.Import("sort", qlsort.Exports)
	qlang.Import("strconv", qlstrconv.Exports)
	qlang.Import("strings", qlstrings.Exports)
	qlang.Import("sync", qlsync.Exports)

	qlang.Import("text_template", qltext_template.Exports)
	qlang.Import("time", qltime.Exports)

	qlang.Import("unicode", qlunicode.Exports)
	qlang.Import("unicode_utf8", qlunicode_utf8.Exports)

	// 3rd party

	qlang.Import("github_topxeq_tk", qlgithubtopxeqtk.Exports)
	qlang.Import("tk", qlgithubtopxeqtk.Exports)
	qlang.Import("github_topxeq_imagetk", qlgithubtopxeqimagetk.Exports)
	qlang.Import("imagetk", qlgithubtopxeqimagetk.Exports)

	qlang.Import("github_beevik_etree", qlgithubbeeviketree.Exports)
	qlang.Import("etree", qlgithubbeeviketree.Exports)
	qlang.Import("github_topxeq_sqltk", qlgithubtopxeqsqltk.Exports)
	qlang.Import("sqltk", qlgithubtopxeqsqltk.Exports)

	qlang.Import("github_topxeq_xmlx", qlgithub_topxeq_xmlx.Exports)

	qlang.Import("github_topxeq_awsapi", qlgithub_topxeq_awsapi.Exports)

	qlang.Import("github_cavaliercoder_grab", qlgithub_cavaliercoder_grab.Exports)

	qlang.Import("github_pterm_pterm", qlgithub_pterm_pterm.Exports)

	qlang.Import("github_domodwyer_mailyak", qlgithub_domodwyer_mailyak.Exports)

	

	qlang.Import("github_fogleman_gg", qlgithub_fogleman_gg.Exports)
	qlang.Import("gg", qlgithub_fogleman_gg.Exports)

	qlang.Import("github_360EntSecGroupSkylar_excelize", qlgithub_360EntSecGroupSkylar_excelize.Exports)

	qlang.Import("github_kbinani_screenshot", qlgithub_kbinani_screenshot.Exports)

	qlang.Import("github_stretchr_objx", qlgithub_stretchr_objx.Exports)

	qlang.Import("github_topxeq_doc2vec_doc2vec", qlgithub_topxeq_doc2vec_doc2vec.Exports)

	qlang.Import("github_aliyun_alibabacloudsdkgo_services_dysmsapi", qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi.Exports)
	qlang.Import("aliyunsms", qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi.Exports)

	// qlang.Import("github_avfs_avfs_fs_memfs", qlgithub_avfs_avfs_fs_memfs.Exports)
	qlang.Import("github_topxeq_afero", qlgithub_topxeq_afero.Exports)
	qlang.Import("memfs", qlgithub_topxeq_afero.Exports)

	qlang.Import("github_topxeq_socks", qlgithub_topxeq_socks.Exports)

	qlang.Import("github_topxeq_regexpx", qlgithub_topxeq_regexpx.Exports)

}

func showHelp() {
	tk.Pl("Gox by TopXeQ V%v\n", versionG)

	tk.Pl("Usage: gox [-v|-h] test.gox, ...\n")
	tk.Pl("or just gox without arguments to start REPL instead.\n")

}

// func compileSource(srcA string) string {
// 	vmT := qlang.New()

// 	tk.Pl("vmT: %v", vmT)

// 	errT := vmT.TXCompile(srcA)

// 	if errT != nil {
// 		return errT.Error()
// 	}

// 	tk.Pl("vmT after: %v", vmT)

// 	return ""

// }

func runInteractiveQlang() int {
	var following bool
	var source string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if following {
			source += "\n"
			fmt.Print("  ")
		} else {
			fmt.Print("> ")
		}

		if !scanner.Scan() {
			break
		}
		source += scanner.Text()
		if source == "" {
			continue
		}
		if source == "quit()" {
			break
		}

		// stmts, err := parser.ParseSrc(source)

		// if e, ok := err.(*parser.Error); ok {
		// 	es := e.Error()
		// 	if strings.HasPrefix(es, "syntax error: unexpected") {
		// 		if strings.HasPrefix(es, "syntax error: unexpected $end,") {
		// 			following = true
		// 			continue
		// 		}
		// 	} else {
		// 		if e.Pos.Column == len(source) && !e.Fatal {
		// 			fmt.Fprintln(os.Stderr, e)
		// 			following = true
		// 			continue
		// 		}
		// 		if e.Error() == "unexpected EOF" {
		// 			following = true
		// 			continue
		// 		}
		// 	}
		// }

		retG = notFoundG

		err := qlVMG.SafeEval(source)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			following = false
			source = ""
			continue
		}

		if retG != notFoundG {
			fmt.Println(retG)
		}

		following = false
		source = ""
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, "ReadString error:", err)
			return 12
		}
	}

	return 0
}

// Non GUI related end



func runFile(argsA ...string) interface{} {
	lenT := len(argsA)

	// full version related start
	
	// full version related end

	if lenT < 1 {
		return nil
	}

	fcT := tk.LoadStringFromFile(argsA[0])

	if tk.IsErrorString(fcT) {
		return tk.Errf("Invalid file content: %v", tk.GetErrorString(fcT))
	}

	return runScript(fcT, "", argsA[1:]...)
}

func runLine(strA string) interface{} {
	argsT, errT := tk.ParseCommandLine(strA)

	if errT != nil {
		return errT
	}

	return runArgs(argsT...)
}

func runArgs(argsA ...string) interface{} {
	argsT := argsA

	if tk.IfSwitchExistsWhole(argsT, "-version") {
		tk.Pl("Gox by TopXeQ V%v", versionG)
		return nil
	}

	if tk.IfSwitchExistsWhole(argsT, "-h") {
		showHelp()
		return nil
	}

	scriptT := tk.GetParameterByIndexWithDefaultValue(argsT, 0, "")

	

	if tk.IfSwitchExistsWhole(argsT, "-initgui") {
		applicationPathT := tk.GetApplicationPath()

		osT := tk.GetOSName()

		if tk.Contains(osT, "inux") {
			tk.Pl("Please visit the following URL to find out how to make Sciter environment ready in Linux: ")

			return nil
		} else if tk.Contains(osT, "arwin") {
			tk.Pl("Please visit the following URL to find out how to make Sciter environment ready in Linux: ")

			return nil
		} else {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/sciter.dll", applicationPathT, "sciter.dll")

			if tk.IsErrorString(rs) {

				return tk.Errf("failed to download Sciter DLL file.")
			}

			tk.Pl("Sciter DLL downloaded to application path.")

			// rs = tk.DownloadFile("http://scripts.frenchfriend.net/pub/webview.dll", applicationPathT, "webview.dll", false)

			// if tk.IsErrorString(rs) {

			// 	return tk.Errf("failed to download webview DLL file.")
			// }

			// rs = tk.DownloadFile("http://scripts.frenchfriend.net/pub/WebView2Loader.dll", applicationPathT, "WebView2Loader.dll", false)

			// if tk.IsErrorString(rs) {

			// 	return tk.Errf("failed to download webview DLL file.")
			// }

			// tk.Pl("webview DLL downloaded to application path.")

			return nil
		}
	}

	ifClipT := tk.IfSwitchExistsWhole(argsT, "-clip")
	ifEmbedT := (codeTextG != "") && (!tk.IfSwitchExistsWhole(argsT, "-noembed"))

	if scriptT == "" && (!ifClipT) && (!ifEmbedT) {

		// autoPathT := filepath.Join(tk.GetApplicationPath(), "auto.gox")
		// autoGxbPathT := filepath.Join(tk.GetApplicationPath(), "auto.gxb")
		autoPathT := "auto.gox"
		autoGxbPathT := "auto.gxb"

		if tk.IfFileExists(autoPathT) {
			scriptT = autoPathT
		} else if tk.IfFileExists(autoGxbPathT) {
			scriptT = autoGxbPathT
		} else {
			initQLVM()

			runInteractiveQlang()

			// tk.Pl("not enough parameters")

			return nil
		}

	}

	encryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-encrypt=", "")

	if encryptCodeT != "" {
		fcT := tk.LoadStringFromFile(scriptT)

		if tk.IsErrorString(fcT) {

			return tk.Errf("failed to load file [%v]: %v", scriptT, tk.GetErrorString(fcT))
		}

		encStrT := tk.EncryptStringByTXDEF(fcT, encryptCodeT)

		if tk.IsErrorString(encStrT) {

			return tk.Errf("failed to encrypt content [%v]: %v", scriptT, tk.GetErrorString(encStrT))
		}

		rsT := tk.SaveStringToFile("//TXDEF#"+encStrT, scriptT+"e")

		if tk.IsErrorString(rsT) {

			return tk.Errf("failed to encrypt file [%v]: %v", scriptT, tk.GetErrorString(rsT))
		}

		return nil
	}

	decryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrypt=", "")

	if decryptCodeT != "" {
		fcT := tk.LoadStringFromFile(scriptT)

		if tk.IsErrorString(fcT) {

			return tk.Errf("failed to load file [%v]: %v", scriptT, tk.GetErrorString(fcT))
		}

		decStrT := tk.DecryptStringByTXDEF(fcT, decryptCodeT)

		if tk.IsErrorString(decStrT) {

			return tk.Errf("failed to decrypt content [%v]: %v", scriptT, tk.GetErrorString(decStrT))
		}

		rsT := tk.SaveStringToFile(decStrT, scriptT+"d")

		if tk.IsErrorString(rsT) {

			return tk.Errf("failed to decrypt file [%v]: %v", scriptT, tk.GetErrorString(rsT))
		}

		return nil
	}

	decryptRunCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrun=", "")

	ifBatchT := tk.IfSwitchExistsWhole(argsT, "-batch")

	if !ifBatchT {
		if tk.EndsWithIgnoreCase(scriptT, ".gxb") {
			ifBatchT = true
		}
	}

	ifBinT := tk.IfSwitchExistsWhole(argsT, "-bin")
	if ifBinT {
	}

	ifExampleT := tk.IfSwitchExistsWhole(argsT, "-example")
	ifGoPathT := tk.IfSwitchExistsWhole(argsT, "-gopath")
	ifLocalT := tk.IfSwitchExistsWhole(argsT, "-local")
	ifAppPathT := tk.IfSwitchExistsWhole(argsT, "-apppath")
	ifRemoteT := tk.IfSwitchExistsWhole(argsT, "-remote")
	ifCloudT := tk.IfSwitchExistsWhole(argsT, "-cloud")
	sshT := tk.GetSwitchWithDefaultValue(argsT, "-ssh=", "")
	ifViewT := tk.IfSwitchExistsWhole(argsT, "-view")
	// ifCompileT := tk.IfSwitchExistsWhole(argsT, "-compile")

	verboseG = tk.IfSwitchExistsWhole(argsT, "-verbose")

	ifMagicT := false
	magicNumberT, errT := tk.StrToIntE(scriptT)

	if errT == nil {
		ifMagicT = true
	}

	var fcT string

	if ifMagicT {
		fcT = getMagic(magicNumberT)

		scriptPathG = ""
	} else if ifExampleT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}
		fcT = tk.DownloadPageUTF8("https://gitee.com/topxeq/gox/raw/master/scripts/"+scriptT, nil, "", 30)

		scriptPathG = ""
	} else if ifRemoteT {
		fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)

		scriptPathG = ""
	} else if ifClipT {
		fcT = tk.GetClipText()

		scriptPathG = ""
	} else if ifEmbedT {
		fcT = codeTextG

		scriptPathG = ""
	} else if ifCloudT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		basePathT, errT := tk.EnsureBasePath("gox")

		gotT := false

		if errT == nil {
			cfgPathT := tk.JoinPath(basePathT, "cloud.cfg")

			cfgStrT := tk.Trim(tk.LoadStringFromFile(cfgPathT))

			if !tk.IsErrorString(cfgStrT) {
				fcT = tk.DownloadPageUTF8(cfgStrT+scriptT, nil, "", 30)

				gotT = true
			}

		}

		if !gotT {
			fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
		}

		scriptPathG = ""
	} else if sshT != "" {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		fcT = downloadStringFromSSH(sshT, scriptT)

		if tk.IsErrorString(fcT) {

			return tk.Errf("failed to get script from SSH: %v", tk.GetErrorString(fcT))
		}

		scriptPathG = ""
	} else if ifGoPathT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		scriptPathG = filepath.Join(tk.GetEnv("GOPATH"), "src", "github.com", "topxeq", "gox", "scripts", scriptT)

		fcT = tk.LoadStringFromFile(scriptPathG)
	} else if ifAppPathT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		scriptPathG = filepath.Join(tk.GetApplicationPath(), scriptT)

		fcT = tk.LoadStringFromFile(scriptPathG)
	} else if ifLocalT {
		if (!tk.EndsWith(scriptT, ".gox")) && (!tk.EndsWith(scriptT, ".ql")) {
			scriptT += ".gox"
		}

		localPathT := getCfgString("localScriptPath.cfg")

		if tk.IsErrorString(localPathT) {
			// tk.Pl("failed to get local path: %v", tk.GetErrorString(localPathT))

			return tk.Errf("failed to get local path: %v", tk.GetErrorString(localPathT))
		}

		// if tk.GetEnv("GOXVERBOSE") == "true" {
		// 	tk.Pl("Try to load script from %v", filepath.Join(localPathT, scriptT))
		// }

		scriptPathG = filepath.Join(localPathT, scriptT)

		fcT = tk.LoadStringFromFile(scriptPathG)
	} else {
		scriptPathG = scriptT
		fcT = tk.LoadStringFromFile(scriptT)
	}

	if tk.IsErrorString(fcT) {
		return tk.Errf("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))
	}

	if tk.StartsWith(fcT, "//TXDEF#") {
		if decryptRunCodeT == "" {
			tk.Prf("Password: ")
			decryptRunCodeT = tk.Trim(tk.GetInputBufferedScan())

			// fcT = fcT[8:]
		}
	}

	if decryptRunCodeT != "" {
		fcT = tk.DecryptStringByTXDEF(fcT, decryptRunCodeT)
	}

	if ifViewT {
		tk.Pl("%v", fcT)

		return nil
	}

	// if ifCompileT {
	// 	initQLVM()

	// 	qlVMG.SetVar("argsG", argsT)

	// 	retG = notFoundG

	// 	endT, errT := qlVMG.SafeCl([]byte(fcT), "")
	// 	if errT != nil {

	// 		// tk.Pl()

	// 		// f, l := qlVMG.Code.Line(qlVMG.Code.Reserve().Next())
	// 		// tk.Pl("Next line: %v, %v", f, l)

	// 		return tk.Errf("failed to compile script(%v) error: %v\n", scriptT, errT)
	// 	}

	// 	tk.Pl("endT: %v", endT)

	// 	errT = qlVMG.DumpEngine()

	// 	if errT != nil {
	// 		return tk.Errf("failed to dump engine: %v\n", errT)
	// 	}

	// 	tk.Plvsr(qlVMG.Cpl.GetCode().Len(), qlVMG.Run())

	// 	return nil
	// }

	if !ifBatchT {
		if tk.RegStartsWith(fcT, `//\s*(GXB|gxb)`) {
			ifBatchT = true
		}
	}

	if ifBatchT {
		listT := tk.SplitLinesRemoveEmpty(fcT)

		// tk.Plv(fcT)
		// tk.Plv(listT)

		for _, v := range listT {
			// tk.Pl("Run line: %#v", v)
			v = tk.Trim(v)

			if tk.StartsWith(v, "//") {
				continue
			}

			rsT := runLine(v)

			if rsT != nil {
				valueT, ok := rsT.(error)

				if ok {
					return valueT
				} else {
					tk.Pl("%v", rsT)
				}
			}

		}

		return nil
	}

	initQLVM()

	qlVMG.SetVar("argsG", argsT)

	retG = notFoundG

	errT = qlVMG.SafeEval(fcT)
	if errT != nil {

		// tk.Pl()

		// f, l := qlVMG.Code.Line(qlVMG.Code.Reserve().Next())
		// tk.Pl("Next line: %v, %v", f, l)

		return tk.Errf("failed to execute script(%v) error: %v\n", scriptT, errT)
	}

	rs, ok := qlVMG.GetVar("outG")

	if ok {
		if rs != nil {
			return rs
		}
	}

	return retG
}

// init the main VM

var retG interface{}
var notFoundG = interface{}(errors.New("not found"))

func initQLVM() {
	if qlVMG == nil {
		qlang.SetOnPop(func(v interface{}) {
			retG = v
		})

		// qlang.SetDumpCode("1")

		importQLNonGUIPackages()

		

		qlVMG = qlang.New()
	}
}

func downloadStringFromSSH(sshA string, filePathA string) string {
	aryT := tk.Split(sshA, ":")

	basePathT, errT := tk.EnsureBasePath("gox")

	if errT != nil {
		return tk.GenerateErrorStringF("failed to find base path: %v", errT)
	}

	if len(aryT) != 5 {
		aryT = tk.Split(tk.LoadStringFromFile(tk.JoinPath(basePathT, "ssh.cfg"))+filePathA, ":")

		if len(aryT) != 5 {
			return tk.ErrStrF("invalid ssh config: %v", "")
		}

	}

	clientT, errT := tk.NewSSHClient(aryT[0], tk.StrToIntWithDefaultValue(aryT[1], 22), aryT[2], aryT[3])

	if errT != nil {
		return tk.ErrToStrF("failed to create SSH client:", errT)
	}

	tmpPathT := tk.JoinPath(basePathT, "tmp")

	errT = tk.EnsureMakeDirsE(tmpPathT)

	if errT != nil {
		return tk.ErrToStrF("failed to create tmp dir:", errT)
	}

	tmpFileT, errT := tk.CreateTempFile(tmpPathT, "")

	if errT != nil {
		return tk.ErrToStrF("failed to create tmp dir:", errT)
	}

	defer os.Remove(tmpFileT)

	errT = clientT.Download(aryT[4], tmpFileT)

	if errT != nil {
		return tk.ErrToStrF("failed to download file:", errT)
	}

	fcT := tk.LoadStringFromFile(tmpFileT)

	return fcT
}

func getCfgString(fileNameA string) string {
	basePathT, errT := tk.EnsureBasePath("gox")

	if errT == nil {
		cfgPathT := tk.JoinPath(basePathT, fileNameA)

		cfgStrT := tk.Trim(tk.LoadStringFromFile(cfgPathT))

		if !tk.IsErrorString(cfgStrT) {
			return cfgStrT
		}

		return tk.ErrStrF("failed to get config string: %v", tk.GetErrorString(cfgStrT))

	}

	return tk.ErrStrF("failed to get config string")
}

var editFileScriptG = `
sciter = github_scitersdk_gosciter
window = github_scitersdk_gosciter_window

htmlT := ` + "`" + `
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<title>Gox Editor</title>
    <style>
    
    plaintext {
      padding:0;
      flow:vertical;
      behavior:plaintext;
      background:#333; border:1px solid #333;
      color:white;
	  overflow:scroll-indicator;
      font-rendering-mode:snap-pixel;
      size:*; 
      tab-size: 4;
    }
    plaintext > text {
      font-family:monospace;
      white-space: pre-wrap;
      background:white;
      color:black;
      margin-left: 3em;
      padding-left: 4dip;
      cursor:text;
      display:list-item;
      list-style-type: index;
      list-marker-color:#aaa;
    }
    plaintext > text:last-child {
      padding-bottom:*;
    }    
    
    plaintext > text:nth-child(10n) {
      list-marker-color:#fff;
    }
    
    
    </style>


	<script type="text/tiscript">
		function colorize() 
		{
			const apply = Selection.applyMark; // shortcut
			const isEditor = this.tag == "plaintext";
			
			// forward declarations:
			var doStyle;
			var doScript;

			// markup colorizer  
			function doMarkup(tz) 
			{
					var bnTagStart = null;
					var tagScript = false;
					var tagScriptType = false;
					var tagStyle = false;
					var textElement;
				
				while(var tt = tz.token()) {
				if( isEditor && tz.element != textElement )       
				{
					textElement = tz.element;
					textElement.attributes["type"] = "markup";
				}
				//stdout.println(tt,tz.attr,tz.value);
				switch(tt) {
					case #TAG-START: {    
						bnTagStart = tz.tokenStart; 
						const tag = tz.tag;
						tagScript = tag == "script";
						tagStyle  = tag == "style";
					} break;
					case #TAG-HEAD-END: {
						apply(bnTagStart,tz.tokenEnd,"tag"); 
						if( tagScript ) { tz.push(#source,"</sc"+"ript>"); doScript(tz, tagScriptType, true); }
						else if( tagStyle ) { tz.push(#source,"</style>"); doStyle(tz, true); }
					} break;
					case #TAG-END:      apply(tz.tokenStart,tz.tokenEnd,"tag"); break;  
					case #TAG-ATTR:     if( tagScript && tz.attr == "type") tagScriptType = tz.value; 
										if( tz.attr == "id" ) apply(tz.tokenStart,tz.tokenEnd,"tag-id"); 
										break;
				}
				}
			}
			
			// script colorizer
			doScript = function(tz, typ, embedded = false) 
			{
				const KEYWORDS = 
				{
				"type"    :true, "function":true, "var"       :true,"if"       :true,
				"else"    :true, "while"   :true, "return"    :true,"for"      :true,
				"break"   :true, "continue":true, "do"        :true,"switch"   :true,
				"case"    :true, "default" :true, "null"      :true,"super"    :true,
				"new"     :true, "try"     :true, "catch"     :true,"finally"  :true,
				"throw"   :true, "typeof"  :true, "instanceof":true,"in"       :true,
				"property":true, "const"   :true, "get"       :true,"set"      :true,
				"include" :true, "like"    :true, "class"     :true,"namespace":true,
				"this"    :true, "assert"  :true, "delete"    :true,"otherwise":true,
				"with"    :true, "__FILE__":true, "__LINE__"  :true,"__TRACE__":true,
				"debug"   :true, "await"   :true 
				};
				
				const LITERALS = { "true": true, "false": true, "null": true, "undefined": true };
				
				var firstElement;
				var lastElement;
			
				while:loop(var tt = tz.token()) {
				var el = tz.element;
				if( !firstElement ) firstElement = el;
				lastElement = el;
				switch(tt) 
				{
					case #NUMBER:       apply(tz.tokenStart,tz.tokenEnd,"number"); break; 
					case #NUMBER-UNIT:  apply(tz.tokenStart,tz.tokenEnd,"number-unit"); break; 
					case #STRING:       apply(tz.tokenStart,tz.tokenEnd,"string"); break;
					case #NAME:         
					{
					var val = tz.value;
					if( val[0] == '#' )
						apply(tz.tokenStart,tz.tokenEnd, "symbol"); 
					else if(KEYWORDS[val]) 
						apply(tz.tokenStart,tz.tokenEnd, "keyword"); 
					else if(LITERALS[val]) 
						apply(tz.tokenStart,tz.tokenEnd, "literal"); 
					break;
					}
					case #COMMENT:      apply(tz.tokenStart,tz.tokenEnd,"comment"); break;
					case #END-OF-ISLAND:  
					// got </scr ipt>
					tz.pop(); //pop tokenizer layer
					break loop;
				}
				}
				if(isEditor && embedded) {
				for( var el = firstElement; el; el = el.next ) {
					el.attributes["type"] = "script";
					if( el == lastElement )
					break;
				}
				}
			};
			
			doStyle = function(tz, embedded = false) 
			{
				const KEYWORDS = 
				{
				"rgb":true, "rgba":true, "url":true, 
				"@import":true, "@media":true, "@set":true, "@const":true
				};
				
				const LITERALS = { "inherit": true };
				
				var firstElement;
				var lastElement;
				
				while:loop(var tt = tz.token()) {
				var el = tz.element;
				if( !firstElement ) firstElement = el;
				lastElement = el;
				switch(tt) 
				{
					case #NUMBER:       apply(tz.tokenStart,tz.tokenEnd,"number"); break; 
					case #NUMBER-UNIT:  apply(tz.tokenStart,tz.tokenEnd,"number-unit"); break; 
					case #STRING:       apply(tz.tokenStart,tz.tokenEnd,"string"); break;
					case #NAME:         
					{
					var val = tz.value;
					if( val[0] == '#' )
						apply(tz.tokenStart,tz.tokenEnd, "symbol"); 
					else if(KEYWORDS[val]) 
						apply(tz.tokenStart,tz.tokenEnd, "keyword"); 
					else if(LITERALS[val]) 
						apply(tz.tokenStart,tz.tokenEnd, "literal"); 
					break;
					}
					case #COMMENT:      apply(tz.tokenStart,tz.tokenEnd,"comment"); break;
					case #END-OF-ISLAND:  
					// got </sc ript>
					tz.pop(); //pop tokenizer layer
					break loop;
				}
				}
				if(isEditor && embedded) {
				for( var el = firstElement; el; el = el.next ) {
					el.attributes["type"] = "style";
					if( el == lastElement )
					break;
				}
				}
			};
			
			var me = this;
			
			function doIt() { 
			
				var typ = me.attributes["type"];

				var syntaxKind = typ like "*html" || typ like "*xml" ? #markup : #source;
				var syntax = typ like "*css"? #style : #script;
				
				var tz = new Tokenizer( me, syntaxKind );
			
				if( syntaxKind == #markup )
				doMarkup(tz);
				else if( syntax == #style )
				doStyle(tz);
				else 
				doScript(tz,typ);
			}
			
			doIt();
			
			// redefine value property
			this[#value] = property(v) {
				get { return this.state.value; }
				set { this.state.value = v; doIt(); }
			};
			
			this.load = function(text,sourceType) 
			{
				this.attributes["type"] = sourceType;
				if( !isEditor )
				text = text.replace(/\r\n/g,"\n"); 
				this.state.value = text; 
				doIt();
			};
			
			this.sourceType = property(v) {
				get { return this.attributes["type"]; }
				set { this.attributes["type"] = v; doIt(); }
			};
			if (isEditor)
					this.on("change", function() {
						this.timer(40ms,doIt);
					});
			

		}
	</script>
	<style>

		@set colorizer < std-plaintext 
		{
			:root { aspect: colorize; }
			
			text { white-space:pre;  display:list-item; list-style-type: index; list-marker-color:#aaa; }
			/*markup*/  
			text::mark(tag) { color: olive; } /*background-color: #f0f0fa;*/
			text::mark(tag-id) { color: red; } /*background-color: #f0f0fa;*/

			/*source*/  
			text::mark(number) { color: brown; }
			text::mark(number-unit) { color: brown; }
			text::mark(string) { color: teal; }
			text::mark(keyword) { color: blue; }
			text::mark(symbol) { color: brown; }
			text::mark(literal) { color: brown; }
			text::mark(comment) { color: green; }
			
			text[type=script] {  background-color: #FFFAF0; }
			text[type=markup] {  background-color: #FFF;  }
			text[type=style]  {  background-color: #FAFFF0; }
		}

		plaintext[type] {
			style-set: colorizer;
		}

		@set element-colorizer 
		{
			:root { 
				aspect: colorize; 
				background-color: #fafaff;
					padding:4dip;
					border:1dip dashed #bbb;
				}
			
			/*markup*/  
			:root::mark(tag) { color: olive; } 
			:root::mark(tag-id) { color: red; }

			/*source*/  
			:root::mark(number) { color: brown; }
			:root::mark(number-unit) { color: brown; }
			:root::mark(string) { color: teal; }
			:root::mark(keyword) { color: blue; }
			:root::mark(symbol) { color: brown; }
			:root::mark(literal) { color: brown; }
			:root::mark(comment) { color: green; }
			}

			pre[type] {
			style-set: element-colorizer;
		}

	</style>
	<script type="text/tiscript">
		// if (view.connectToInspector) {
		// 	view.connectToInspector(rootElement, inspectorIpAddress);
		// }

		//stdout.println("__FOLDER__:", __FOLDER__);
		//stdout.println("__FILE__:", __FILE__);

		function isErrStr(strA) {
			if (strA.substr(0, 6) == "TXERROR:") {
				return true;
			}

			return false;
		}

		function getErrStr(strA) {
			if (strA.substr(0, 6) == "TXERROR:") {
				return strA.substr(6);
			}

			return strA;
		}

		function getConfirm(titelA, msgA) {
			var result = view.msgbox { 
				type:#question,
				title: titelA,
				content: msgA, 
				//buttons:[#yes,#no]
				buttons: [
					{id:#yes,text:"Ok",role:"default-button"},
					{id:#cancel,text:"Cancel",role:"cancel-button"}]                               
				};

			return result;
		}

		function showInfo(titelA, msgA) {
			var result = view.msgbox { 
				type:#information,
				title: titelA,
				content: msgA, 
				//buttons:[#yes,#no]
				buttons: [
					{id:#cancel,text:"Close",role:"cancel-button"}]                               
				};

			return result;
		}

		function showError(titelA, msgA) {
			var result = view.msgbox { 
				type:#alert,
				title: titelA,
				content: msgA, 
				//buttons:[#yes,#no]
				buttons: [
					{id:#cancel,text:"Close",role:"cancel-button"}]                               
				};

			return result;
		}

		function getScreenWH() {
//			view.prints(String.printf("screenBoxO: %v, %v", 1, 2));
			var (w, h) = view.screenBox(#frame, #dimension)
//			view.prints(String.printf("screenBox: %v, %v", w, h));

			view.move((w-800)/2, (h-600)/2, 800, 600);

			return String.printf("%v|%v", w, h);
		}

		var editFileNameG = "";
		var editFileCleanFlagG = "";

		function updateFileName() {
			$(#fileNameLabelID).html = (editFileNameG + editFileCleanFlagG);
		}

		function selectFileJS() {
			//var fn = view.selectFile(#open, "Gotx Files (*.gt,*.go)|*.gt;*.go|All Files (*.*)|*.*" , "gotx" );
			var fn = view.selectFile(#open);
			view.prints(String.printf("fn: %v", fn));
			//view.prints(String.printf("screenBox: %v", view.screenBox(#frame, #dimension)));

			if (fn == undefined) {
				return;
			}

			var fileNameT = URL.toPath(fn);

			var rs = view.loadStringFromFile(fileNameT);

			if (isErrStr(rs)) {
				showError("Error", String.printf("Failed to load file content: %v", getErrStr(rs)));
				return;
			}

			$(plaintext).attributes["type"] = "text/script";

			$(plaintext).value = rs;

			editFileNameG = fileNameT;

			editFileCleanFlagG = "";

			updateFileName();

			// return fn;
		}

		function editFileLoadClick() {
			if (editFileCleanFlagG != "") {
			
				var rs = getConfirm("Please confirm", "File modified, load another file anyway?");

				if (rs != #yes) {
					return;
				}

			}

			selectFileJS();
		}

		function editFileSaveAsClick() {
			var fn = view.selectFile(#save);
			view.prints(String.printf("fn: %v", fn));

			if (fn == undefined) {
				return;
			}

			var fileNameT = URL.toPath(fn);

			var textT = $(plaintext).value;

			var rs = view.saveStringToFile(textT, fileNameT);

			if (isErrStr(rs)) {
				showError("Error", String.printf("Failed to save file content: %v", getErrStr(rs)));
				return;
			}

			editFileNameG = fileNameT;
			editFileCleanFlagG = "";
			updateFileName();

			showInfo("Info", "Saved.");

		}

		function editFileSaveClick() {
			if (editFileNameG == "") {
				editFileSaveAsClick();

				return;
			}

			var textT = $(plaintext).value;

			var rs = view.saveStringToFile(textT, editFileNameG);

			if (isErrStr(rs)) {
				showError("Error", String.printf("Failed to save file content: %v", getErrStr(rs)));
				return;
			}

			editFileCleanFlagG = "";
			updateFileName();

			showInfo("Info", "Saved.");
		}

		function editRunClick() {
			view.close();
			// view.exit();
		}

		function getInput(msgA) {
			var res = view.dialog({ 
				html: ` + "`+\"`\"+`" + `
				<html>
				<body>
				  <center>
					  <div style="margin-top: 10px; margin-bottom: 10px;">
						  <span>` + "`+\"`\"+`" + `+msgA+` + "`+\"`\"+`" + `</span>
					  </div>
					  <div style="margin-top: 10px; margin-bottom: 10px;">
						  <input id="mainTextID" type="text" />
					  </div>
					  <div style="margin-top: 10px; margin-bottom: 10px;">
						  <input id="submitButtonID" type="button" value="Ok" />
						  <input id="cancelButtonID" type="button" value="Cancel" />
					  </div>
				  </center>
				  <script type="text/tiscript">
					  $(#submitButtonID).onClick = function() {
						  view.close($(#mainTextID).value);
					  };
  
					  $(#cancelButtonID).onClick = function() {
						  view.close();
					  };
				  </scr` + "`+\"`\"+`" + `+` + "`+\"`\"+`" + `ipt>
				</body>
				</html>
				` + "`+\"`\"+`" + `
			  });
  
			  return res;
		  }

		event click $(#btnEncrypt)
		{
		  	var res = getInput("Secure Code");

			if (res == undefined) {
				return;
			}

			var sourceT = $(plaintext).value;

			var encStrT = view.encryptText(sourceT, res);
		
			if (isErrStr(encStrT)) {
				showError("Error", String.printf("failed to encrypt content: %v",getErrStr(encStrT)));
				return;
			}
		
			$(plaintext).value = "\/\/TXDEF#" + encStrT;
			editFileCleanFlagG = "*";
			updateFileName();
		}
	
		event click $(#btnDecrypt)
		{
		  	var res = getInput("Secure Code");

			if (res == undefined) {
				return;
			}

			var sourceT = $(plaintext).value;

			var encStrT = view.decryptText(sourceT, res);
		
			if (isErrStr(encStrT)) {
				showError("Error", String.printf("failed to decrypt content: %v",getErrStr(encStrT)));
				return;
			}
		
			$(plaintext).value = encStrT;
			editFileCleanFlagG = "*";
			updateFileName();
		}
	
		event click $(#btnRun)
		{
			var res = getInput("Arguments to pass to script")

			if (res == undefined) {
				return;
			}

			var rs = view.runScript($(plaintext).value, res);

			showInfo("Result", rs)
		 
		  	// view.prints(String.printf("result = %v", rs));
		}
	

		function editCloseClick() {
			view.close();
			// view.exit();
		}

		function editFile(fileNameA) {
			var fcT string;

			//view.prints("fileNameA: "+fileNameA);

			if (fileNameA == "") {
				editFileNameG = "";

				fcT = "";

				editFileCleanFlagG = "*";
			} else {
				editFileNameG = fileNameA;

				fcT = view.loadStringFromFile(fileNameA);

//		if tk.IsErrorString(fcT) {
//			tk.Pl("failed to load file %v: %v", editFileNameG, tk.GetErrorString(fcT))
//			return

//		}

				editFileCleanFlagG = "";
			}

			//view.prints(fcT);

			$(plaintext).attributes["type"] = "text/script";

			$(plaintext).value = fcT;

			updateFileName();

		}

		function self.ready() {

			//$(plaintext).value = "<html>\n<body>\n<span>abc</span>\n</body></html>";

			$(#btnLoad).onClick = editFileLoadClick;
			$(#btnSave).onClick = editFileSaveClick;
			$(#btnSaveAs).onClick = editFileSaveAsClick;
			// $(#btnEncrypt).onClick = editFEncryptClick;
			// $(#btnDecrypt).onClick = editDecryptClick;
			// $(#btnRun).onClick = editRunClick;
			$(#btnClose).onClick = editCloseClick;

			$(plaintext#source).onControlEvent = function(evt) {
				switch (evt.type) {
					case Event.EDIT_VALUE_CHANGED:      
						editFileCleanFlagG = "*";
						updateFileName();
						return true;
				}
			};

		}
	</script>

</head>
<body>
	<div style="margin-top: 10px; margin-bottom: 10px;"><span id="fileNameLabelID"></span></div>
	<div style="margin-top: 10px; margin-bottom: 10px;">
		<button id="btn1" style="display: none">Load...</button>
		<button id="btnLoad">Load</button>
		<button id="btnSave">Save</button>
		<button id="btnSaveAs">SaveAs</button>
		<button id="btnEncrypt">Encrypt</button>
		<button id="btnDecrypt">Decrypt</button>
		<button id="btnRun">Run</button>
		<button id="btnClose">Close</button>
	</div>
	<plaintext#source type="text/html" style="font-size: 1.2em;"></plaintext>

</body>
</html>
` + "`" + `

// htmlT = tk.LoadStringFromFile(tk.JoinPath(path_filepath.Dir(scriptPathG), "editFileSciter.st"))

// tk.CheckErrorString(htmlT)

runtime.LockOSThread()

w, err := window.New(sciter.DefaultWindowCreateFlag, sciter.DefaultRect)

checkError(err)

w.SetOption(sciter.SCITER_SET_SCRIPT_RUNTIME_FEATURES, sciter.ALLOW_FILE_IO | sciter.ALLOW_SOCKET_IO | sciter.ALLOW_EVAL | sciter.ALLOW_SYSINFO)

w.LoadHtml(htmlT, "")

w.SetTitle("Gox Editor")

w.DefineFunction("prints", func(args) {
	tk.Pl("%v", args[0].String())
	return sciter.NewValue("")
})

w.DefineFunction("loadStringFromFile", func(args) {
	rs := tk.LoadStringFromFile(args[0].String())
	return sciter.NewValue(rs)
})

w.DefineFunction("saveStringToFile", func(args) {
	rs := tk.SaveStringToFile(args[0].String(), args[1].String())
	return sciter.NewValue(rs)
})

w.DefineFunction("encryptText", func(args) {
	rs := tk.EncryptStringByTXDEF(args[0].String(), args[1].String())
	return sciter.NewValue(rs)
})

w.DefineFunction("decryptText", func(args) {
	rs := tk.DecryptStringByTXDEF(args[0].String(), args[1].String())
	return sciter.NewValue(rs)
})

w.DefineFunction("runScript", func(args) {
	rs := runScript(args[0].String(), "", args[1].String())
	return sciter.NewValue(tk.Spr("%v", rs))
})

w.DefineFunction("exit", func(args) {
	os.Exit(1);
})

data, _ := w.Call("getScreenWH") //, sciter.NewValue(10), sciter.NewValue(20))
// fmt.Println("data:", data.String())

fileNameT := tk.GetParameterByIndexWithDefaultValue(argsG, 0, "")

w.Call("editFile", sciter.NewValue(fileNameT))

w.Show()

// screenshot = github_kbinani_screenshot

// tk.Plvsr(screenshot.NumActiveDisplays(), screenshot.GetDisplayBounds(0))

// bounds := screenshot.GetDisplayBounds(0)

// img, err := screenshot.CaptureRect(bounds)
// if err != nil {
// 	panic(err)
// }
// fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
// file, _ := os.Create(fileName)

// image_png.Encode(file, img)

// file.Close()

// fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)

w.Run()

`

func editFile(fileNameA string, argsA ...string) {
	rs := runScriptX(editFileScriptG, argsA...)

	if rs != notFoundG {
		// tk.Pl("%v", rs)
	}

}

func main() {
	// var errT error

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Exception: ", err)
		}
	}()

	test()

	rand.Seed(time.Now().Unix())

	rs := runArgs(os.Args[1:]...)

	if rs != nil {
		valueT, ok := rs.(error)

		if ok {
			if valueT != spec.Undefined && valueT != notFoundG {
				tk.Pl("Error: %T %v", valueT, valueT)
			}
		} else {
			tk.Pl("%v", rs)
		}
	}

}

func test() {

}

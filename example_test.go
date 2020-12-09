package golib_os_test
import (
	"testing"
	myos "github.com/weizhouBlue/golib_os"
	"fmt"
	"time"
	"os"
)

//====================================

func Test_seatch(t *testing.T){

	myos.EnableLog=true


	executable:="ls"
	if path , err:=myos.SearchExecutable( executable ); err!= nil {
		fmt.Println(  "failed to find "+ executable )
		t.FailNow()
	}else{
		fmt.Println( executable + " is @ " + path )

	}


}


func Test_simple_simpleCmd(t *testing.T){

	myos.EnableLog=true

	// exec command
	cmd:="echo $WELAN ; echo aaa "
	// addtional environment 
	env:=[]string{
		"WELAN=12345",
		"TOM=uit",
	}
	// stdin for cmd
	stdin_msg:="this is stdin msg"
	// o for no auto timeout
	timeout_second:=5
	chanFinish , _ , chanStdoutMsg , chanStderrMsg , chanErr , exitedCode,  e  :=myos.RunCmd( cmd, env , stdin_msg , timeout_second )
	if e!=nil {
		fmt.Println(  "failed to exec "+ cmd )
		t.FailNow()

	}
	<-chanFinish

	if data , ok := <-chanErr ; ok {
		// return code with no succeed
		fmt.Println(  "err : "+ data )
	}else{
		//fmt.Println("ok for cmd"  )

		if data , ok := <-chanStdoutMsg ; ok {
			fmt.Println(  "stdoutMsg: "+ data )
		}
		if data , ok := <-chanStderrMsg ; ok {
			fmt.Println(  "stderrMsg: "+ data )
		}
		if data , ok := <-exitedCode ; ok {
			fmt.Println(  "exitedCode: ",  data )
		}

	}

}

func Test_simple_longCmd(t *testing.T){

	myos.EnableLog=true

	// exec command
	cmd:="sleep 10d"
	// addtional environment 
	env:=[]string{
		"WELAN=12345",
		"TOM=uit",
	}
	// stdin for cmd
	stdin_msg:="this is stdin msg"
	// o for no auto timeout
	timeout_second:=0
	chanFinish , chanCancel , chanStdoutMsg , chanStderrMsg , chanErr ,_ , e  :=myos.RunCmd( cmd, env , stdin_msg , timeout_second )
	if e!=nil {
		fmt.Println(  "failed to exec "+ cmd )
		t.FailNow()

	}

	//wait for cmd
	select{
	case <- chanFinish:
		fmt.Println("cmd finish ")
	case <- time.After(5*time.Second) : 
		fmt.Println("cmd timeout , cancel it")
		close(chanCancel)
	}
	
	//read msg
	if data , ok := <-chanErr ; ok {
		// return code with no succeed
		fmt.Println(  "err : "+ data )
	}else{
		//fmt.Println("ok for cmd"  )

		if data , ok := <-chanStdoutMsg ; ok {
			fmt.Println(  "stdoutMsg: "+ data )
		}
		if data , ok := <-chanStderrMsg ; ok {
			fmt.Println(  "stderrMsg: "+ data )
		}
	}

}






func Test_json1(t *testing.T){

	data:=map[string] string {
		"k1": "v1" ,
		"k2": "v2" ,

	}
	filePath:="./json_test"

	if e:=myos.WriteJsonToFile( filePath , data ) ; e!=nil {
		fmt.Println(  "failed to WriteJsonToFile " )
		t.FailNow()
	}


	if jsondata , e:=myos.ReadJsonFromFile( filePath ) ; e!=nil {
		fmt.Println(  "failed to ReadJsonFromFile " )
		t.FailNow()
	}else{
		fmt.Printf(  "json data: %v \n" , string(jsondata) )
	}

	os.Remove(filePath)


	
	myos.EmptyFile(filePath)

	size , _ := myos.FileSize( filePath )
	fmt.Printf("size1eddd : %v \n "  , size  ) 




}


func Test_path(t *testing.T){
	fmt.Println( myos.GetMyExecName() )
	fmt.Println( myos.GetMyExecDir() )
	fmt.Println( myos.GetMyRunDir() )
	
	
}

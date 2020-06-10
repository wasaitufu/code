package sftp

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	azip "github.com/alexmullins/zip"
)

func DownLoadBigZip() {
	client := NewSftpClient("yzh_prod", "TGZUkaJ7WdyQ2vUG", "49.4.70.156", "", 13550)
	err := client.Connect()
	if err != nil {
		fmt.Errorf("client.Connect failed, err=%v", err)
		return
	}
	defer client.Close()
	fmt.Println("client.Connect success")

	// 21006530000000810000000000000000409461475955381196.zip
	// 21006530000000810000000000000000408319486287352074.zip
	file, err := client.Open("yzh_prod/21006530000000810000000000000000408319486287352074.zip")
	if err != nil {
		fmt.Errorf("client.Open failed, err=%v", err)
		return
	}
	defer file.Close()
	fmt.Println("client.Open success")

	// ***相比较ioutil.ReadAll，创建临时文件的方法可以节省大部分内存消耗
	tmpFile, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	// io.CopyBuffer()
	io.Copy(tmpFile, file)

	zipReadCloser, err := azip.OpenReader(tmpFile.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range zipReadCloser.File {
		func() {
			fmt.Println(f.Name)
			// 解压密码
			if f.IsEncrypted() {
				f.SetPassword("247578")
			}
		}()
	}

	// 用缓冲流读取文件会更快，下面代码还有待完成
	//var bt []byte
	//const BufferSize = 1024 * 1024 * 10
	//buffer := make([]byte, BufferSize)
	//for {
	//	_, err := file.Read(buffer)
	//	if err != nil {
	//		if err != io.EOF {
	//			fmt.Println(err)
	//		}
	//		break
	//	}
	//	bt = BytesCombine(bt, buffer)
	//}

	// 读取解压文件流，一个加密压缩包中多个无加密压缩包
	//zipReader, err := azip.NewReader(bytes.NewReader(bt), int64(len(bt)))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//for _, f := range zipReader.File {
	//	func() {
	//		fmt.Println(f.Name)
	//		// 解压密码
	//		if f.IsEncrypted() {
	//			f.SetPassword("247578")
	//		}
	//
	//		//file, errLocal := ioutil.TempFile("", "tempfile")
	//		//if errLocal != nil {
	//		//	fmt.Printf("ioutil.TempFile failed, err=%v", errLocal)
	//		//}
	//
	//		//io.Copy(file, f.zipr)
	//
	//		_, errLocal := zip.OpenReader(f.FileInfo().Name())
	//		// 文件仍是压缩包，因此需要再以zip形式读取
	//		//zReader, errLocal := zip.NewReader(bytes.NewReader(fileBits), int64(len(fileBits)))
	//		if errLocal != nil {
	//			fmt.Println("zip.NewReader failed, err=%+v", errLocal)
	//			return
	//		}
	//
	//		////// 打开文件
	//		//fileRc, errLocal := f.Open()
	//		//if errLocal != nil {
	//		//	fmt.Errorf("file.Open failed, err=%+v", err)
	//		//	return
	//		//}
	//		//defer fileRc.Close()
	//		//
	//		//// 读取文件流
	//		//if errLocal != nil {
	//		//	fmt.Errorf("ioutil.ReadAll failed, err=%+v", err)
	//		//	return
	//		//}
	//
	//		//// 文件仍是压缩包，因此需要再以zip形式读取
	//		//zReader, errLocal := zip.NewReader(bytes.NewReader(fileBits), int64(len(fileBits)))
	//		//fileBits, errLocal := ioutil.ReadAll(fileRc)
	//		//if errLocal != nil {
	//		//	ctx.Logger().Errorf("zip.NewReader failed, err=%+v", errLocal)
	//		//	return
	//		//}
	//		//
	//		//// 遍历文件
	//		//for _, f := range zReader.File {
	//		//	// 优雅退出监听
	//		//	select {
	//		//	case <-ctx.Done():
	//		//		ctx.Logger().Errorf("HandleReceipt ctx quit, bankAccountNo=%s, date=%s, err=%+v", param.BaseParam.BankAccountNo, param.BaseParam.Date, ctx.Err())
	//		//		// TODO: 通知回单服务进行重试
	//		//		return
	//		//	default:
	//		//	}
	//		//
	//		//	// 并发入队
	//		//	handleConChan <- byte('t')
	//		//
	//		//	fcp := f
	//		//	par.Go(func() {
	//		//		handleReceiptUploadAndPush(ctx, fcp, ossBucket, tikac)
	//		//		// 并发出队
	//		//		<-handleConChan
	//		//	})
	//		//}
	//	}()
	//}
}

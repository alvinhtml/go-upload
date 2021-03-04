package file

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// exists 判断文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Upload 上传文件
func Upload(c *gin.Context) {

	md5 := c.Param("md5")
	index := c.Param("index")

	log.Println("md5", md5)

	file, err := c.FormFile("file")

	if err != nil {
		log.Println("ERROR: upload file failed. ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("ERROR: upload file failed. %s", err),
		})
	}
	dir := fmt.Sprintf(`/Users/alvin/tmp/uploads/` + md5)
	path := fmt.Sprintf(dir + "/part-" + index)
	// path := fmt.Sprintf(`/Users/alvin/tmp/uploads/` + file.Filename)

	if !exists(dir) {
		err := os.MkdirAll(dir, 0766)
		if err != nil {
			log.Println("ERROR: mkdir failed. ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": fmt.Sprintf("ERROR: mkdir failed. %s", err),
			})
		}
	}

	// 保存文件至指定路径
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		log.Println("ERROR: save file failed. ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("ERROR: save file failed. %s", err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":      "file upload succ.",
			"filepath": path,
		})
	}
}

// Merge 合并文件
func Merge(c *gin.Context) {
	// md5 := c.Param("md5")
	// chunkTotal := c.Param("chunkTotal") // 文件分块总数
	// total, err := strconv.ParseInt(chunkTotal, 10, 64)
	// if err != nil {
	// 	log.Println("ERROR: Chunk total is not a valid number. ", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"msg": fmt.Sprintf("ERROR: Chunk total is not a valid number. %s", err),
	// 	})
	// }
	// f, err := os.Open(fileName)
	// if err != nil {
	//     fmt.Println("can't opened this file")
	//     return err
	// }
	// defer f.Close()
	// s := make([]byte, 4096)
}

// ExecMerge 用exec命令合并文件
func ExecMerge(c *gin.Context) {
	md5 := c.Param("md5")

	fileDir := fmt.Sprintf(`/Users/alvin/tmp/uploads/` + md5)
	cmdString := fmt.Sprintf("/bin/cat " + fileDir + "/part-* > " + fileDir + "/" + md5 + ".tar")

	cmd := exec.Command("/bin/bash", "-c", cmdString)
	err := cmd.Run()

	if err != nil {
		log.Println("ERROR: Merge file failed. ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("ERROR: Merge file failed. %s", err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file merge succ.",
		})
	}
}

// TestCors 测试 CORS
func TestCors(c *gin.Context) {
	fmt.Print(1)
	c.JSON(http.StatusOK, gin.H{
		"cors": "ok!",
	})
}

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

//パッケージ修飾子を_にすることでエクスポートされた名前は見えない
func init() {
	fmt.Println("init()")
	//sql.Openはデータベースへの接続を行わない（直観に反して）
	//データストアへのコネクションは必要になった時に初めて遅延評価される
	//取得したDBオブジェクトがアクセス可能なものかどうか確認したければdb.Ping()を使用する
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/world")
	/*
		data sourceについて
		rootの場合理由はよくわからないが使用するデータベースをソースの文字列に記載しても
		選択されず、そのままselectすると「1046: No database selected」となる。
		ただしuse hoge;というクエリを実行しておけば問題ない。
	*/
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	file, err := os.Create("cityNames.txt")
	if err != nil {
		log.Fatal(err)
	}

	if err := getCityNames(file); err != nil {
		log.Fatal(err)
	}
	sum, err := getSumCityPopulation()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("各都市の合計人口は%v人\n", sum)
}

func getSumCityPopulation() (int64, error) {

	var (
		sum int64
	)
	rows, err := db.Query("select SUM(Population) from city")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		//結果を詰め込む変数名は関係ないっぽい。型は関係ある
		err := rows.Scan(&sum)
		if err != nil {
			continue
		}
	}
	return sum, rows.Err()
}

func getCityNames(w io.WriteCloser) error {
	defer w.Close()
	var (
		name string
	)
	rows, err := db.Query("select Name from city")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		//結果を詰め込む変数名は関係ないっぽい。型は関係ある
		err := rows.Scan(&name)
		if err != nil {
			continue
		}
		fmt.Fprintf(w, "とれた：name=%v\n", name)
	}
	return rows.Err()
}

func sumPopulations() (int64, error) {
	f, err := os.Open("pops.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var sum int64
	scanner := bufio.NewScanner(f)

	fmt.Println(f.Name())
	for scanner.Scan() {
		population, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println(err)
			break
		}
		sum += int64(population)
	}
	return sum, scanner.Err()
}

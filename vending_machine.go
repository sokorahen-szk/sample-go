package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	appName = "自販機"
	ps      = "> "
)

type Coin int

func NewCoin(value int) (*Coin, error) {
	if !(value == 1 || value == 5 || value == 10 || value == 50 || value == 100 || value == 500) {
		return nil, errors.New("硬貨じゃない？")
	}

	ins := Coin(value)
	return &ins, nil
}
func (c Coin) Int() int       { return int(c) }
func (c Coin) String() string { return fmt.Sprintf("%d", c.Int()) }

type Bill int

func NewBill(value int) (*Bill, error) {
	err := errors.New("紙幣じゃないよね？")
	if !(value == 1000 || value == 2000 || value == 5000 || value == 10000) {
		return nil, err
	}

	if value%1000 != 0 {
		return nil, err
	}

	ins := Bill(value)
	return &ins, nil
}
func (b Bill) Int() int       { return int(b) }
func (b Bill) String() string { return fmt.Sprintf("%d", b.Int()) }

type DrinkName string
type Drink struct {
	name  DrinkName
	price int
}

const (
	OrangeJuice DrinkName = "オレンジジュース"
)

func NewDrink(name DrinkName, price int) *Drink {
	return &Drink{
		name:  name,
		price: price,
	}
}

func (i Drink) GetPrice() int {
	return i.price
}

func (i Drink) GetName() string {
	return string(i.name)
}

var (
	coins  []*Coin
	bills  []*Bill
	drinks []*Drink
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	d := []*Drink{
		NewDrink(OrangeJuice, 120),
	}
	drinks = append(drinks, d...)

	for {
		fmt.Print(ps)
		scanner.Scan()
		act := scanner.Text()

		var err error
		switch act {
		case "into_coins":
			fmt.Print("硬貨を入れてね：")
			err = intoMoney("coins")
			break
		case "into_bills":
			fmt.Print("紙幣を入れてね：")
			err = intoMoney("bills")
			break
		case "money":
			fmt.Println(fmt.Sprintf("預かり金：%d 円", SumBill()+SumCoin()))
			break
		case "help":
			fmt.Println(fmt.Sprintf(`%s の使い方
into_coins
　硬貨を入力する。使える硬貨：1円,5円,10円,50円,100円

into_bills
　紙幣を入力する。使える紙幣：1000円,2000円,5000円,10000円

money
　預かり金を確認する。

e
　%s 終了する。

help
　ヘルプを表示する。
`, appName, appName))
		}

		if err != nil {
			fmt.Println(err.Error())
		}

		if act == "e" {
			break
		}
	}
}

func intoMoney(input string) error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in := scanner.Text()

	num, _ := strconv.Atoi(in)

	if input == "coins" {
		res, err := convertToCoin(num)
		if err != nil {
			return err
		}
		coins = append(coins, res...)
	} else if input == "bills" {
		res, err := convertToBill(num)
		if err != nil {
			return err
		}
		bills = append(bills, res...)
	}

	return nil
}

func convertToCoin(_coins int) ([]*Coin, error) {
	if _coins < 1 {
		return nil, nil
	}
	coinFormats := []int{
		500, 100, 50, 10, 5, 1,
	}
	resources := make([]*Coin, 0)
	errBit := false
	for {
		for i, coinFormat := range coinFormats {
			if _coins%coinFormat == 0 {
				coin, err := NewCoin(coinFormat)
				if err != nil {
					return nil, err
				}
				resources = append(resources, coin)
				_coins -= coinFormat
				break
			}
			if i+1 == len(coinFormats) {
				errBit = true
			}
		}
		if errBit {
			return nil, errors.New("与えられた硬貨に変なコインが含まれてない？")
		}
		if 1 > _coins {
			break
		}
	}
	return resources, nil
}

func convertToBill(_bills int) ([]*Bill, error) {
	if _bills < 1 {
		return nil, nil
	}
	billFormats := []int{
		10000, 5000, 2000, 1000,
	}
	resources := make([]*Bill, 0)
	errBit := false
	for {
		for i, billFormat := range billFormats {
			if _bills%billFormat == 0 {
				coin, err := NewBill(billFormat)
				if err != nil {
					return nil, err
				}
				resources = append(resources, coin)
				_bills -= billFormat
				break
			}
			if i+1 == len(billFormats) {
				errBit = true
			}
		}
		if errBit {
			return nil, errors.New("与えられた硬貨に変なコインが含まれてない？")
		}
		if 1 > _bills {
			break
		}
	}

	return resources, nil
}

func SumCoin() int {
	sum := 0
	for _, coin := range coins {
		sum += coin.Int()
	}
	return sum
}

func SumBill() int {
	sum := 0
	for _, bill := range bills {
		sum += bill.Int()
	}
	return sum
}

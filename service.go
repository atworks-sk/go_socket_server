package main

import (
	"fmt"
	"log"
	"strconv"
)

func Service0001(s string) string {
	product := s[0:4]
	sex := s[4:5]
	age := s[5:6]
	log.Println("product : " + product)
	log.Println("sex : " + sex)
	log.Println("age : " + age)

	premiumOrigin := 0
	discountRate := 0

	// 상품
	if product == "a001" {
		premiumOrigin = 100000
	} else if product == "b001" {
		premiumOrigin = 200000
	} else if product == "c001" {
		premiumOrigin = 300000
	}

	// 성별
	if sex == "F" {
		discountRate += 10
	}

	// 나이
	switch age {
	case "2":
		discountRate += 20
	case "3":
		discountRate += 15
	case "4":
		discountRate += 10
	case "5":
		discountRate += 5
	default:
		discountRate += 0
	}

	log.Println(strconv.Itoa(premiumOrigin))
	log.Println(strconv.Itoa(discountRate))
	//Math.round(premium * (100f-discountRate) / 100f);
	premium := premiumOrigin * (100 - discountRate) / 100
	log.Println(strconv.Itoa(premium))
	return fmt.Sprintf("%8d%2d%8d", premium, discountRate, premiumOrigin)
}
func Service0002(s string) string {
	product := s[0:4]
	text := ""
	if product == "a001" {
		text = "a001 불라불라"
	} else if product == "b001" {
		text = "b001 불리불라"
	} else if product == "c001" {
		text = "c001 불리불라"
	} else {
		text = "상품코드가 잘못되었습니다"
	}
	return fmt.Sprintf("%1000s", text)

	// fmt.Sprintf("%s%s%09d%s", start_time, service_id, second, result)
}

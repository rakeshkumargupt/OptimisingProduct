package main

import (
	"fmt"	
	"io/ioutil"
	"encoding/json"
	"os"
)

type Product []struct {
	ImageURLs []string `json:"imageURLs"`
	ProductTitle string `json:"productTitle"`
	Price int `json:"price"`
	DiscountedPrice int `json:"discountedPrice"`
	DiscountPer string `json:"discountPer"`
	MarketplaceName string `json:"marketplaceName"`
	AffliateLink string `json:"affliateLink"`
	AvailableColor []string `json:"availableColor"`
	AvailableSize []string `json:"availableSize"`
	ProductID string `json:"productId"`
	Productfeature []string `json:"productfeature"`
	ProductAttribute struct {
		SuitableFor string `json:"Suitable For"`
		Neck string `json:"Neck"`
	} `json:"productAttribute"`
	Rating string `json:"rating"`
	Reviews int `json:"reviews"`
	RecommendedProduct []int `json:"recommendedProduct"`
}

type UserPref []struct {
	UserID string `json:"userId"`
	UserName string `json:"userName"`
	UserGeo string `json:"userGeo"`
	PreferredSize []string `json:"preferredSize"`
	PreferredColor []string `json:"preferredColor"`
	PreferredStyle []string `json:"preferredStyle"`
	UserInterest []string `json:"userInterest"`
	UserAttribute struct {
		Age string `json:"age"`
		Height string `json:"height"`
		BodyType string `json:"bodyType"`
		Color string `json:"color"`
		Gender string `json:"gender"`
	} `json:"userAttribute"`
}

type ArrayDataType []string
var UserMap map[string]ArrayDataType
var ColorMap, SizeMap map[int]string

func main() {

	// Reading Product JSON File
	ProductJSONFile, PJErr  := os.Open("Product.json")
	if PJErr != nil {
		fmt.Println("Error Opening Product JSON File : ", PJErr)
		return
	}
	defer ProductJSONFile.Close()
	ProductJSONData, PJDataErr  := ioutil.ReadAll(ProductJSONFile)
	if PJDataErr != nil {
		fmt.Println("Error reading ProductJSONData :", PJDataErr)
	}
	//	fmt.Println(string(ProductJSONData))
	var ProductVar Product
	json.Unmarshal(ProductJSONData, &ProductVar)
	//	fmt.Println(ProductVar)


	//Reading UserPref JSON File
	UserPrefJSONFile, UPJErr := os.Open("UserPref.json")
	if UPJErr != nil {
		fmt.Println("Error Opening UserPref JSON FIle", UPJErr)
		return
	}

	defer UserPrefJSONFile.Close()

	UserPrefJSONData, UPJDataErr := ioutil.ReadAll(UserPrefJSONFile)
	if UPJDataErr != nil {
		fmt.Printf("Error reading UserPrefJSONData :", UPJDataErr)
		return
	}
	var UserPrefVar UserPref
	json.Unmarshal(UserPrefJSONData, &UserPrefVar)
	//fmt.Println(UserPrefVar)

	// Getting ProductTitle Array
		//ProductTitleArray := ProductVar[0].ProductTitle
		//fmt.Println(ProductTitleArray)

	// Preparing array of Attribute from


	//Getting User Preference Attribute
	for userIndex := range UserPrefVar{

		UserId := UserPrefVar[userIndex].UserID
		PreferredStyleData := UserPrefVar[userIndex].PreferredStyle
		PreferredColorData := UserPrefVar[userIndex].PreferredColor
		PreferredSizeData := UserPrefVar[userIndex].PreferredSize
		UserInterestData := UserPrefVar[userIndex].UserInterest

		fmt.Println(UserId)
		//fmt.Println(PreferredColorData)
		//fmt.Println(PreferredSizeData)
		//fmt.Println(PreferredStyleData)
		//fmt.Println(UserInterestData)

		UserMap = make(map[string]ArrayDataType)
		UserMap["PreferredColor"] = PreferredColorData
		UserMap["PreferredStyle"] = PreferredStyleData
		UserMap["PreferredSize"] = PreferredSizeData
		UserMap["UserInterest"] = UserInterestData

		//for userMapIndex := range UserMap{
		//	fmt.Println(" ", userMapIndex, ": ", UserMap[userMapIndex])
		//}

		for productIndex := range ProductVar{

			ProductId := ProductVar[productIndex].ProductID
			AvailableColorData := ProductVar[productIndex].AvailableColor
			AvailableSizeData := ProductVar[productIndex].AvailableSize
			Productfeaturedata := ProductVar[productIndex].Productfeature

			fmt.Println(ProductId)
			fmt.Println(AvailableColorData)
			fmt.Println(AvailableSizeData)
			fmt.Println(Productfeaturedata)

			ColorMap = make(map[int]string)
			for colorIndex := range AvailableColorData{
				ColorMap[colorIndex] = AvailableColorData[colorIndex]
			}
			fmt.Println(ColorMap)

			SizeMap = make(map[int]string)
			for sizeIndex := range AvailableSizeData{
				SizeMap[sizeIndex] = AvailableSizeData[sizeIndex]
			}
			fmt.Println(SizeMap)

			SuitableForData := ProductVar[productIndex].ProductAttribute.SuitableFor
			fmt.Println(SuitableForData)



		}


	}




	//fmt.Println(ProductVar[0].AvailableColor[0])



//		json.Unmarshal([]byte(c), &data)
//	     fmt.Println(data)
   //      fmt.Println(data.imageURLs[0])


















}

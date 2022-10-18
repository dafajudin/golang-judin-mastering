package controllers

import (
	"go-hero/models"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func GetAllHero(c echo.Context) error {
	result, err := models.GetAllHero()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func PostHero(c echo.Context) error {

	//accept data from client
	nama_hero 		:= c.FormValue("nama_hero")
	role 			:= c.FormValue("role")
	emblem 			:= c.FormValue("emblem")
	battle_spell 	:= c.FormValue("battle_spell")

	//call function PostHero from models
	result, err := models.PostHero(nama_hero, role, emblem, battle_spell)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateHero(c echo.Context) error {

	//accept id from client
	id 				:= c.FormValue("id")
	nama_hero 		:= c.FormValue("nama_hero")
	role 			:= c.FormValue("role")
	emblem 			:= c.FormValue("emblem")
	battle_spell 	:= c.FormValue("battle_spell")

	//convert id to int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	//call function UpdateHero from models
	result, err := models.UpdateHero(conv_id, nama_hero, role, emblem, battle_spell)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
	
}

func DeleteHero(c echo.Context) error{
	//accept id from client
	id := c.FormValue("id")

	//convert id to int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	//call function DeleteHero from models
	result, err := models.DeleteHero(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

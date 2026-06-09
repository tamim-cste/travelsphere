package api

import (
	"encoding/json"
	"travelsphere/services"
	"travelsphere/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// WishlistAPIController handles /api/wishlist CRUD
type WishlistAPIController struct {
	beego.Controller
}

// Get returns all wishlist items as JSON
func (w *WishlistAPIController) Get() {
	items := services.GetWishlist()
	utils.SendSuccess(&w.Controller, items)
}

// Post creates a new wishlist item
func (w *WishlistAPIController) Post() {
	var body struct {
		CountryName string `json:"country_name"`
		Note        string `json:"note"`
		Status      string `json:"status"`
	}
	if err := json.Unmarshal(w.Ctx.Input.RequestBody, &body); err != nil {
		utils.SendError(&w.Controller, 400, "Invalid request body")
		return
	}

	item, err := services.AddToWishlist(body.CountryName, body.Note, body.Status)
	if err != nil {
		utils.SendError(&w.Controller, 400, err.Error())
		return
	}

	w.Ctx.ResponseWriter.WriteHeader(201)
	w.Data["json"] = map[string]interface{}{"success": true, "data": item}
	w.ServeJSON()
}

// Put updates note/status of an existing item
func (w *WishlistAPIController) Put() {
	id := w.Ctx.Input.Param(":id")

	var body struct {
		Note   string `json:"note"`
		Status string `json:"status"`
	}
	if err := json.Unmarshal(w.Ctx.Input.RequestBody, &body); err != nil {
		utils.SendError(&w.Controller, 400, "Invalid request body")
		return
	}

	item, err := services.UpdateWishlistItem(id, body.Note, body.Status)
	if err != nil {
		utils.SendError(&w.Controller, 404, err.Error())
		return
	}
	utils.SendSuccess(&w.Controller, item)
}

// Delete removes a wishlist item
func (w *WishlistAPIController) Delete() {
	id := w.Ctx.Input.Param(":id")
	if err := services.DeleteWishlistItem(id); err != nil {
		utils.SendError(&w.Controller, 404, err.Error())
		return
	}
	utils.SendSuccess(&w.Controller, map[string]string{"message": "deleted"})
}

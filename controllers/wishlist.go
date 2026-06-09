package controllers

import "travelsphere/services"


type WishlistController struct {
	BaseController
}

func (w *WishlistController) Prepare() {
	w.BaseController.Prepare()
	loggedIn, ok := w.Data["LoggedIn"].(bool)
	if !ok || !loggedIn {
		w.Redirect("/login", 302)
	}
}

func (w *WishlistController) Get() {
	items := services.GetWishlist()
	w.Data["Title"] = "Travel Wishlist"
	w.Data["WishlistItems"] = items
	w.Layout = "layout/main.tpl"
	w.TplName = "wishlist.tpl"
}

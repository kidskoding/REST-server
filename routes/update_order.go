package routes

import (
	"REST-server/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method. Use PUT", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	orderID := vars["id"]
	if orderID == "" {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	id, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var updatedOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	collection := models.GetCollection()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedOrder}

	result, err := collection.UpdateOne(r.Context(), filter, update)
	if err != nil {
		http.Error(w, "Failed to update order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bson.M{
		"message": "Order updated successfully",
	})
}

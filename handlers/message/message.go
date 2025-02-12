package message

import (
	"log"
	"mygoapp/libs/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TikTokChatMessage struct {
	ID             primitive.ObjectID     `bson:"_id,omitempty" json:"_id"`
	Emotes         []string               `bson:"emotes" json:"emotes"`
	Comment        string                 `bson:"comment" json:"comment"`
	UserID         string                 `bson:"userId" json:"userId"`
	UniqueID       string                 `bson:"uniqueId" json:"uniqueId"`
	UserSceneTypes []string               `bson:"userSceneTypes" json:"userSceneTypes"`
	UserDetails    map[string]interface{} `bson:"userDetails" json:"userDetails"`
}

func GetMessages(c *gin.Context) {
	items := c.Query("items")
	// Convertir items a un número entero usando strconv.Atoi
	num, err := strconv.Atoi(items)
	if err != nil {
		// Si hay un error al convertir, puedes manejarlo de alguna manera,
		// por ejemplo, responder con un código de estado HTTP y un mensaje.
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "La cadena 'items' no es un número válido",
		})
		return
	}

	ctx, client, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	databaseName := "chatbot"
	collectionName := "tiktokchatmessages"

	db := client.Database(databaseName)
	collection := db.Collection(collectionName)

	defer client.Disconnect(ctx)

	findOptions := options.Find().SetLimit(int64(num)) // Opcional: limita el número de documentos devueltos
	cur, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var documents []TikTokChatMessage

	// Imprime los datos de la colección
	for cur.Next(ctx) {
		var doc TikTokChatMessage
		err := cur.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		documents = append(documents, doc)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, documents)
}

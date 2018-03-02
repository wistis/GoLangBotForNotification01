package main

import (

	"log"
	"net/http"
	"github.com/googollee/go-socket.io"
	"github.com/julienschmidt/httprouter"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"time"
)

var dbconn_str  = "user:pasword@tcp(37.46.129.210:3306)/dbname"
var newzakaz = 0   //Новый заказ  0- нету иначе ID заказа


type customServer struct {
	Server *socketio.Server
}


func (s *customServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	s.Server.ServeHTTP(w, r)
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	router := httprouter.New()

	router.GET("/addzakaz", sendzakaz) //Роутер уведомления о заказ

	server.On("connection", func(so socketio.Socket) {

		log.Println("on connection")

		so.Join("chat")

		go robot(so); //запуск слушателя новых заказов
		so.On("chat message", func(msg string) {

			log.Println("emit:", so.Emit("chat message", msg));
			so.BroadcastTo("chat", "chat message", msg);
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})

		server.On("error", func(so socketio.Socket, err error) {
			log.Println("error:", err)
		})

	});
	wsServer := new(customServer)
	wsServer.Server = server
	http.Handle("/socket.io/", wsServer)

	http.Handle("/", router)
	router.NotFound = http.FileServer(http.Dir("asset"))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
func robot(so socketio.Socket) {
	i := 1
	klk := 0;
	for i < 2 {

		time.Sleep(time.Millisecond * 5000)
		if (newzakaz > 0) {
			klk = newzakaz;
			updatedb(klk);
			newzakaz = 0;
			so.BroadcastTo("chat", "chat message", klk);
			so.Emit("chat message", klk);
		}

	}

}
func sendzakaz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = provLoginIndb();
}

func updatedb(newzakaz int) {
	db, err := sql.Open("mysql", dbconn_str)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	_ = db.QueryRow("UPDATE   clients SET no_view=1 WHERE id = ? ", newzakaz)  //установка статуса заказа в базе
	newzakaz = 0;
	log.Println("База обновлена");
	// defer the close till after the main function has finished
	// executing
	defer db.Close();
}

func provLoginIndb() (int) {
	db, err := sql.Open("mysql", dbconn_str)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	var id int;
	row := db.QueryRow("SELECT id FROM clients WHERE no_view = 0  ORDER BY id DESC") //выбор заказа
	_ = row.Scan(&id)

	newzakaz = id;
	log.Println(newzakaz);
	defer db.Close();
	return id
}

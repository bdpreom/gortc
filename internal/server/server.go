package server

import "flag" 

var (
	addr = flag.String("addr,":" " ",os.Getenv("PORT"),"")
	cert = flag.String("cert","","") 
	key = flag.String("key","","")
)


func Run() error {
	flag.Parse()
	if *addr == ":"{
		*addr = ":8080"
	}

	engine := html.New("./views",".html")
	app := fibre.New(fiber.Config{views: engine})	
	app.Use(logger.New())
	app.Use(cors.New())



	app.Get("/",handlers.Welcome)
	app.Get("/room/create",handlers.RoomCreate)
	app.Get("/room/:uuid",handlers.Room)
	app.Get("/room/:uuid/websocket",)
	app.Get("/room/:uuid/chat",handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket",websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket",websocket.New(handlers.RoomViewerWebsocket))
	app.Get ("/strea/:ssuid", handlers.Stream)
	app.Get ("/ stream/:ssuid/websocket",)
	app.Get ("/stream/ssuid/chat/websocket",)
	app.Get("/stream/:ssuid/viewer/websocket")

  











}
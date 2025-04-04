BookStore –  книжного магазин. Оно помогает управлять книгами, авторами и пользователями. Приложение разделено на несколько частей, чтобы его было легче поддерживать и развивать.

Как устроен проект:
 internal/db – отвечает за подключение к базе данных и хранение данных.
 internal/delivery – принимает запросы от пользователей (например, "покажи все книги") и отправляет ответы.
 internal/models – здесь описаны структуры данных, например, как выглядит "Книга" (название, автор, цена).
 internal/repository – здесь происходят операции с базой данных: добавление, удаление, обновление книг.
 internal/routes – задает, какие запросы куда направлять, например, /books – список книг, /authors – список авторов.
 internal/service – обрабатывает логику приложения, например, проверяет, можно ли купить книгу.
 
 PostManmen teksergende mende localhost 8080 


 
	r.GET("/books", bookHandler.GetAllBooks)- tolyk kitaptardy korsetedi
	r.GET("/books/:id", bookHandler.GetBookByID)- Id arkylly izdeidi
	r.POST("/books", bookHandler.CreateBook)- jana kitap kosady
	r.PUT("/books/:id", bookHandler.UpdateBook)-  kitapty janartady
	r.DELETE("/books/:id", bookHandler.DeleteBook)- kitapty joyadi




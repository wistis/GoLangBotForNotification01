# GoLangBotOrders
При поступлении звонка на сайт или обращения с сайта . 
Записываются данные в базу данных.Отсылается запрос боту методом GET
http://localhost:5000/addzakaz

На сайте с CRM (index.html - пример получения уведомлений от бота ), работает бот написанный на GOLANG . 
К которому подключается клиент(index.html)по SOCKET IO.
 Если менеджер в кабинете или открыта вкладка с CRM то бот посылает менеджеру в браузер уведомление о новом заказе.
В этот момент у менеджера вслывает окно о новом заказе или чтото еще чтото

 
# GoLangBotOrders
When you receive a call to the site or access to the site.
The data is written to the database. A bot request is sent to the GET method
http: // localhost: 5000 / addzakaz

On the site with CRM (index.html - an example of receiving notifications from the bot), the bot written on GOLANG works.
To which the client is connected (index.html) by SOCKET IO.
  If the manager is in the cabinet or the tab with the CRM is opened, then the bot sends a notification to the manager to the browser about the new order.
At this point, the manager has a window about a new order popping up or something else chtoto

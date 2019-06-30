# LogEvent (Логирование событий)

Написать функцию логирования <code>LogEvent</code>, на вход которой приходят 
события типа <code>HwAccepted</code> (домашняя работа принята) и <code>HwSubmitted</code> 
(студент отправил дз) для этого - спроектировать и реализовать интерфейс <code>Event</code>. 
Для события <code>HwAccepted</code> мы хотим логировать дату, айди и грейд, 
для <code>HwSubmitter</code> - дату, id и комментарий, например:

    2019-01-01 submitted 3456 "please take a look at my homework"
    2019-01-01 accepted 3456 4

Псевдокод:

    type HwAccepted struct {
        Id int
        Grade int
    }
    
    type HwSubmitted struct {
        Id int
        Code string
        Comment string
    }
    
    type Event interface  {
    }

    func LogEvent(e Event, w io.Writer) {
    }
 
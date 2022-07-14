Приветствую.

Самое простое что пришло в голову это генерация блоков на основе unixtime текущего.
Предположил что можно фаерволом/балансером ограничить коннекты до 1 раз в сек (Тз не противоречит)
Конечно это все только как Proof of concept. Минусов хватает: генерировать блоки может дольше чем 1 секунда, 
избыточность памяти , удаление блока(нарушение цепочки)

Реализация сервера пошаговая и прямолинейная для экономии времени. В реальности сделал бы через паттерн loop.
Пример https://github.com/edelars/dsn/blob/main/pkg/api/ws/handler/get_devices_status.go

Сервер одноклиентный, опять же для экономии времени да и тз это не противоречит
(Надо добавить прослойку между blockchain и server для хранения раздельных данных по клиентам  как минимум)

Hash выбрал так как простая реализация и можно использовать GPU ускорение (в теории)

Greetings.

The simplest thing that came to mind is the generation of blocks based on the current unixtime.
I suggested that it is possible to limit connections to 1 per second with a firewall / balancer (no info in task)
Of course, this is all only as a Proof of concept. There are enough minuses: it can generate blocks longer than 1 second,
memory redundancy, deleting a block (breaking the chain)

Server implementation is step-by-step and straightforward to save time. In reality, I would do it with loop pattern.
Example https://github.com/edelars/dsn/blob/main/pkg/api/ws/handler/get_devices_status.go

The server is single-client, again, to save time.(No info in task)
(It is necessary to add a layer between the blockchain and the server to store separate data on clients at least)

Hash was chosen because it is a simple implementation and you can use GPU acceleration (in theory)

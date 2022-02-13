/* Описание проекта.
В этом проекте мы займемся обработкой заказов клиентов. Программа будет поддерживать заказы National и International.
Поведение обоих этих интерфейсов будет опираться на абстракцию интерфейса.
Для проекта мы используем модули Go, и я предполагаю, что вы обладаете базовыми знаниями их работы. Если же нет, то это не критично.
Определение функции main будет простым, так как мы импортируем пакет Order и просто вызовем из него функцию New. Этот пакет, в свою очередь, будет содержать логику примера:
package main
*/
package main

import (
	"interfaces/order"
)

func main() {
	order.New() // мы просто импортируем пакет и вызываем в нем функцию New().
}

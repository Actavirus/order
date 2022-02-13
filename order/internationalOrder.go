/* Этот файл является первой фактической реализацией интерфейса Operations.
Сначала мы создали тип структуры InternationalOrder, определив с помощью структуры Order его свойства и объекты.
Далее идет функция инициализации NewInternationalOrder, которая будет устанавливать товары для заказа, информацию о клиенте и адрес доставки.
Для инициализации ProductDetail, Client и ShippingAddress мы используем вспомогательную функцию, которую вскоре тоже реализуем.
В оставшейся части файла мы объявляем фактическую реализацию функций FillOrderSummary, Notify и PrintOrderDetails.
Теперь можно сказать, что тип структуры InternationalOrder реализует интерфейс Operations, потому что содержит определения всех его методов.
Круто!
*/
package order

import "fmt"

// структура international
var international = &InternationalOrder{}

// структура InternationalOrder
type InternationalOrder struct {
	Order
}

// фукнция NewInternationalOrder
func NewInternationalOrder() *InternationalOrder {
	international.products = append(international.products, GetProductDetail("Lap Top", 450, 1, 450.50))
	international.products = append(international.products, GetProductDetail("Video Game", 600, 2, 1200.50))
	international.Client = SetClient("Carl", "Smith", "carlsmith@gmail.com", "9658521365")
	international.ShippingAddress = SetShippingAddress("Colfax Avenue", "Seattle", "USA", "45712")
	return international
}

// функция FillOrderSummary
func (into *InternationalOrder) FillOrderSummary() {
	var extraFee float32 = 0.5                  // дополнительная плата
	var taxes float32 = 0.25                    // налог
	var shippingCost float32 = 35               // стоимость доставки
	subtotal = CalculateSubTotal(into.products) // передаём в функцию CalculateSubTotal данные по всем продуктам и поучаем общую стоимость.

	totalBeforeTax = (subtotal + shippingCost)
	totalTaxes = (taxes * subtotal)
	totalExtraFee = (totalTaxes * extraFee)
	total = (subtotal + totalTaxes) + totalExtraFee
	into.Summary = Summary{
		total:          total,
		subtotal:       subtotal,
		totalBeforeTax: totalBeforeTax,
	}
}

// функция Notify
func (into *InternationalOrder) Notify() {
	email := into.Client.email
	name := into.Client.name
	phone := into.Client.phone

	fmt.Println()
	fmt.Println("---Международный заказ---")
	fmt.Println("Уведомление: ", name)
	fmt.Println("Отправить уведомление на электронную почту :", email)
	fmt.Println("Отправить СМС-уведомление на номер ", phone)
	fmt.Println("Отправить уведомление на WhatsApp:", phone)
}

// функция PrintOrderDetails - метод вывода информации о международном заказе
func (into *InternationalOrder) PrintOrderDetails() {
	fmt.Println()
	fmt.Println("Международный заказ")
	fmt.Println("Информация о заказе: ")
	fmt.Println("-- Наименование товара: ", into.Order.Summary) // пытаемся указать наименвание товара
	fmt.Println("-- ИТОГО до налогооблажения: ", into.Summary.totalBeforeTax)
	fmt.Println("-- Промежуточный итог: ", into.Summary.subtotal)
	fmt.Println("-- Всего: ", into.Summary.total)
	fmt.Printf("-- Адрес доставки: %s %s %s \n", into.ShippingAddress.street, into.ShippingAddress.city, into.ShippingAddress.country)
	fmt.Printf("-- Клиент: %s %s \n", into.Client.name, into.Client.lastName)
}

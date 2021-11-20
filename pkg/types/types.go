package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Status представляет собой статус платежа
type Category string

//Status представляет собой статус платежа
type Status string

// Предопределённые статусы платежей
  const (
    StatusOk Status="OK"
    StatusFail Status="FAIL"
    StatusInProgress Status="INPROGRESS"
  )

//   Payment представляет информавцию о платеже
type Payment struct {
 ID int
 Amount Money
 Category Category
 Status Status
}
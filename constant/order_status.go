package constant

const (
	OrderStatusOrderCreated          = "order_created"
	OrderStatusWaitingForPayment     = "waiting_for_payment"
	OrderStatusPaymentSuccess        = "payment_success"
	OrderStatusPaymentFailed         = "payment_failed"
	OrderStatusOrderCanceled         = "order_cancelled"
	OrderStatusOrderVoided           = "order_voided"
	OrderStatusOrderVoidedWithRefund = "order_voided_with_refund"
	OrderStatusRefundRequested       = "refund_requested"
)

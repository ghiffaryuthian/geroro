package service

import (
  "github.com/bukalapak/bigmom/database"
  "github.com/bukalapak/bigmom/entity"
  "github.com/bukalapak/bigmom/utility"
)

func GetFeedback(currentUser entity.BukalapakTokenResourceOwner, transactionID int64) (entity.Feedback, entity.CustomError) {
  db, _ := database.InitMysql()
  defer db.Close()

  var transaction entity.Transaction
  var feedback entity.Feedback

  db.QueryRowx("SELECT * FROM transactions WHERE id=?", transactionID).StructScan(&transaction)

  if transaction.ID == 0 {
    return feedback, entity.TransactionNotFoundError
  }

  if currentUser.ID != transaction.BuyerID && currentUser.ID != transaction.SellerID && !utility.IsAllowedUser(currentUser) {
    return feedback, entity.UserUnauthorizedError
  }

  db.QueryRowx("SELECT * FROM feedbacks WHERE transaction_id=?", transactionID).StructScan(&feedback)

  if feedback.ID == 0 {
    if transaction.Remarkable() {
      return feedback, entity.FeedbackNotFoundError
    } else {
      feedback = BuildAutoFeedback(&transaction)
    }
  }

  return feedback, entity.CustomError{}
}

func BuildAutoFeedback(t *entity.Transaction) entity.Feedback {
  return entity.Feedback{
    Body:     "Feedback otomatis untuk pembayaran " + t.Description,
    Positive: true,
  }
}

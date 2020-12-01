# Golang SQL fast toolbox
### Named parameters, Transactions and Context

- Faster *named parameters*, no reflection!
- Simpler consistency - *Transactions* always commit or rollback! 
- SQL with *Context*, you should prevent edge problems.
- Improving Developer experience enforncing Good practices.

## Named parameters 

```go
    FetchRows(db, ctx, "SELECT me FROM devs WHERE skills = :skill AND yeart = :year;", Named{"skill": skill, "year": 30}
```
This fix: `https://github.com/go-sql-driver/mysql/issues/561`

## Single or Multiple Transaction in stashed environment 

```go
// main.go
c := mysql.ConfigFromEnvs("TEST") 

db, err := mysql.Open(c, ctx)
if err != nil {
      // do smtg clever 
}
//db.SetMaxLifetimeMins(15) optional set it
```

```go
// in repo.go
// ctx = context.Background() or reuse the ctx from the request/response framework
tx, err := mysql.TxCreate(db, ctx) 
if mysql.IsErrorRollback(err, tx) {
    //do smtg clever, tx is already rollbacked
}

// start using the transaction

incrementalOId, err := mysql.TxPush(tx, ctx, "insert into `order` (name) values (:name)", mysql.Named{"name": "value to insert"})
if mysql.IsErrorRollback(err, tx) {
    // do smtg clever is already rollbacked
}

incrementalIId, err := mysql.TxPush(tx, ctx, "insert into `item` (order_id) values (:order_id)", mysql.Named{"order_id": incrementalOId})
if mysql.IsErrorRollback(err, tx) {
    // do smtg clever is already rollbacked
}

if tx.Commit() != nil {
	 // do smtg clever
}
```

 
  
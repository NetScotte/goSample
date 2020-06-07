package mysql

import "testing"

func TestQueryRowById(t *testing.T) {
	dbtype := "mysql"
	dsn := "root:123lloi@tcp(192.168.199.219:3306)/golang_db"
	err := initDB(dbtype, dsn)
	if err != nil {
		t.Errorf("Failed initDB: %v\n", err)
		return
	}
	user, err := queryRowById(1)
	if err != nil {
		t.Errorf("Failed queryRowById: %v\n", err)
		return
	}
	t.Logf("get user: %+v", user)
}

func TestQueryData(t *testing.T) {
	dbtype := "mysql"
	dsn := "root:123lloi@tcp(192.168.199.218:3306)/golang_db"
	err := initDB(dbtype, dsn)
	if err != nil {
		t.Errorf("Failed initDB: %v\n", err)
	}
	err = queryData(4)
	if err != nil {
		t.Errorf("Failed queryData: %v\n", err)
	}
}

func TestInsertData(t *testing.T) {
	dbtype := "mysql"
	dsn := "root:123lloi@tcp(192.168.199.218:3306)/golang_db"
	err := initDB(dbtype, dsn)
	if err != nil {
		t.Errorf("Failed initDB: %v\n", err)
	}
	err = insertData("freeze", 22)
	if err != nil {
		t.Errorf("Failed insertData: %v\n", err)
	}
}

func TestUpdateData(t *testing.T) {
	dbtype := "mysql"
	dsn := "root:123lloi@tcp(192.168.199.218:3306)/golang_db"
	err := initDB(dbtype, dsn)
	if err != nil {
		t.Errorf("Failed initDB: %v\n", err)
	}
	err = updateData("netliu", 18)
	if err != nil {
		t.Errorf("Failed updateData: %v\n", err)
	}
}

func TestDeleteData(t *testing.T) {
	dbtype := "mysql"
	dsn := "root:123lloi@tcp(192.168.199.218:3306)/golang_db"
	err := initDB(dbtype, dsn)
	if err != nil {
		t.Errorf("Failed initDB: %v\n", err)
	}
	err = deleteData(3)
	if err != nil {
		t.Errorf("Failed deleteData: %v\n", err)
	}
}

func TestPrepareQueryData(t *testing.T) {
	dbtype := "mysql"
	dsn := "root:123lloi@tcp(192.168.199.218:3306)/golang_db"
	err := initDB(dbtype, dsn)
	if err != nil {
		t.Errorf("Failed initDB: %v\n", err)
	}
	err = prepareQueryData(1)
	if err != nil {
		t.Errorf("Failed prepareQueryData: %v\n", err)
	}
}

func TestTransaction(t *testing.T) {
	dbtype := "mysql"
	dsn := "root:123lloi@tcp(192.168.199.219:3306)/golang_db"
	err := initDB(dbtype, dsn)
	if err != nil {
		t.Errorf("Failed initDB: %v\n", err)
		return
	}
	err = transaction()
	if err != nil {
		t.Errorf("Failed transaction: %v\n", err)
		return
	}
}



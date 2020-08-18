package persistance

import (
	"github.com/dgraph-io/badger"
	"github.com/sirupsen/logrus"
)

const (
	DefaultPath = "./db/blocks"

	KeyLastHash = "last_hash"
)

type Persistance struct {
	db *badger.DB
}

type Serializable interface {
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
}

func (p *Persistance) Init(path string, block Serializable, hash []byte) ([]byte, error) {
	logrus.Info("Initializing the database...")
	dbOptions := badger.DefaultOptions(path)
	db, err := badger.Open(dbOptions)
	if err != nil {
		logrus.Error("Failed to initialize database: ", db, err)
		return nil, err
	}
	p.db = db

	var lastHash []byte

	err = p.db.Update(func(transaction *badger.Txn) error {
		lastHashKey := KeyLastHash
		item, err := transaction.Get([]byte(lastHashKey))

		if err != nil {
			if err == badger.ErrKeyNotFound {
				// We have never saved this, meaning this is the first time
				// Therefore we need to create the DB
				logrus.Warn("No previous blockchain found. Creating a new one...")
				serialData, err := block.Serialize()
				if err != nil {
					return err
				}
				if err := transaction.Set(hash, serialData); err != nil {
					return err
				}
				if err := transaction.Set([]byte(KeyLastHash), hash); err != nil {
					return err
				}
				lastHash = hash
				return nil
			}
			return err
		}

		value, err := item.ValueCopy(nil)
		if err != nil {
			logrus.Error("Error occured while getting item value by key: ", KeyLastHash, item, err)
			return err
		}
		lastHash = value
		return nil

	})
	if err != nil {
		logrus.Error("Failed to run Init transaction in the database: ", err)
		return nil, err
	}

	return lastHash, nil
}

func (p *Persistance) Get(key []byte) ([]byte, error) {
	var value []byte
	err := p.db.View(func(transaction *badger.Txn) error {
		item, err := transaction.Get(key)
		if err != nil {
			logrus.Error("Error occured while getting item by key: ", key, item, err)
			return err
		}
		value, err = item.ValueCopy(nil)
		if err != nil {
			logrus.Error("Error occured while getting item value by key: ", key, item, err)
			return err
		}
		return err
	})
	if err != nil {
		logrus.Error("Failed to run Get transaction in the database: ", err)
		return nil, err
	}
	return value, nil
}

func (p *Persistance) GetLastHash() ([]byte, error) {
	logrus.Info("Getting last hash from the database...")
	return p.Get([]byte(KeyLastHash))
}

func (p *Persistance) SaveBlock(hash []byte, block Serializable) error {
	err := p.db.Update(func(transaction *badger.Txn) error {
		serialData, err := block.Serialize()
		if err != nil {
			return err
		}
		if err := transaction.Set(hash, serialData); err != nil {
			return err
		}
		if err := transaction.Set([]byte(KeyLastHash), hash); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		logrus.Error("Failed to run SaveBlock transaction in the database: ", err)
		return err
	}
	return nil
}

func (p *Persistance) Iterate(prefix []byte, block Serializable, callback func(value []byte) error) error {
	err := p.db.View(func(transaction *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		iterator := transaction.NewIterator(opts)
		defer iterator.Close()

		for iterator.Seek(prefix); iterator.ValidForPrefix(prefix); iterator.Next() {
			item := iterator.Item()
			err := item.Value(callback)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

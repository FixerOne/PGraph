package config

import (
	"github.com/pkg/errors"
)

// database code. Need to map to the database code (DataStoreConfig) in the configuration yaml file.
const (
	SQLDB     string = "sqldb"
	COUCHDB   string = "couch"
	CACHEGRPC string = "cacheGrpc"
	USERGRPC  string = "userGrpc"
)

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file.
// Client app use those to retrieve use case from the container
const (
	REGISTRATION string = "registration"
	LISTUSER     string = "listUser"
	LISTCOURSE   string = "listCourse"
)

// data service code. Need to map to the data service code (DataConfig) in the configuration yaml file.
const (
	USERDATA   string = "userData"
	CACHEDATA  string = "cacheData"
	TXDATA     string = "txData"
	COURSEDATA string = "courseData"
)

func validateConfig(appConfig AppConfig) error {
	err := validateDataStore(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateLogger(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	useCase := appConfig.UseCase
	err = validateUseCase(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateLogger(appConfig AppConfig) error {
	zc := appConfig.ZapConfig
	key := zc.Code
	zcMsg := " in validateLogger doesn't match key = "
	if ZAP != key {
		errMsg := ZAP + zcMsg + key
		return errors.New(errMsg)
	}
	lc := appConfig.LorusConfig
	key = lc.Code
	if LOGRUS != lc.Code {
		errMsg := LOGRUS + zcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateDataStore(appConfig AppConfig) error {
	sc := appConfig.SQLConfig
	key := sc.Code
	scMsg := " in validateDataStore doesn't match key = "
	if SQLDB != key {
		errMsg := SQLDB + scMsg + key
		return errors.New(errMsg)
	}
	cc := appConfig.CouchdbConfig
	key = cc.Code
	if COUCHDB != key {
		errMsg := COUCHDB + scMsg + key
		return errors.New(errMsg)
	}
	cgc := appConfig.CacheGrpcConfig
	key = cgc.Code
	if CACHEGRPC != key {
		errMsg := CACHEGRPC + scMsg + key
		return errors.New(errMsg)
	}

	ugc := appConfig.UserGrpcConfig
	key = ugc.Code
	if USERGRPC != key {
		errMsg := USERGRPC + scMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateUseCase(useCase UseCaseConfig) error {
	err := validateRegistration(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateListCourse(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateListUser(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateRegistration(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.Registration
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if REGISTRATION != key {
		errMsg := REGISTRATION + rcMsg + key
		return errors.New(errMsg)
	}
	key = rc.UserDataConfig.Code
	if USERDATA != key {
		errMsg := USERDATA + rcMsg + key
		return errors.New(errMsg)
	}
	key = rc.TxDataConfig.Code
	if TXDATA != key {
		errMsg := TXDATA + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateListUser(useCaseConfig UseCaseConfig) error {
	lc := useCaseConfig.ListUser
	key := lc.Code
	luMsg := " in validateListUser doesn't match key = "
	if LISTUSER != key {
		errMsg := LISTUSER + luMsg + key
		return errors.New(errMsg)
	}
	key = lc.CacheDataConfig.Code
	if CACHEDATA != key {
		errMsg := CACHEDATA + luMsg + key
		return errors.New(errMsg)
	}
	return nil
}
func validateListCourse(useCaseConfig UseCaseConfig) error {
	lc := useCaseConfig.ListCourse
	key := lc.Code
	lcMsg := " in validateListCourse doesn't match key = "
	if LISTCOURSE != key {
		errMsg := LISTCOURSE + lcMsg + key
		return errors.New(errMsg)
	}
	key = lc.CourseDataConfig.Code
	if COURSEDATA != key {
		errMsg := COURSEDATA + lcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

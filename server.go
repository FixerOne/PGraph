package main

import (
	handlerBasequestion "pgraph/handler/basequestion"
	handlerBasesection "pgraph/handler/basesection"
	handlerBaseteststypes "pgraph/handler/baseteststypes"
	handlerBasic "pgraph/handler/basic"
	handlerCompany "pgraph/handler/company"
	handlerDocumentstype "pgraph/handler/documentstype"
	handlerProject "pgraph/handler/project"
	handlerProtocol "pgraph/handler/protocol"
	handlerTest "pgraph/handler/test"
	handlerTestquestion "pgraph/handler/testquestion"
	handlerTestType "pgraph/handler/testtype"
	handlerUser "pgraph/handler/user"

	"github.com/gin-gonic/gin"
)

var ()

func main() {

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	basicHandler := handlerBasic.New(server)
	basicHandler.StartHandlers()

	companyHandler := handlerCompany.New(server)
	companyHandler.StartHandlers()

	userHandler := handlerUser.New(server)
	userHandler.StartHandlers()

	projectHandler := handlerProject.New(server)
	projectHandler.StartHandlers()

	testHandler := handlerTest.New(server)
	testHandler.StartHandlers()

	protocolHandler := handlerProtocol.New(server)
	protocolHandler.StartHandlers()

	testtypeHandler := handlerTestType.New(server)
	testtypeHandler.StartHandlers()

	basequestionHandler := handlerBasequestion.New(server)
	basequestionHandler.StartHandlers()

	basesectionHandler := handlerBasesection.New(server)
	basesectionHandler.StartHandlers()

	baseteststypesHandler := handlerBaseteststypes.New(server)
	baseteststypesHandler.StartHandlers()

	documentstypeHandler := handlerDocumentstype.New(server)
	documentstypeHandler.StartHandlers()

	testquestionHandler := handlerTestquestion.New(server)
	testquestionHandler.StartHandlers()

	server.Run(":8686")

}

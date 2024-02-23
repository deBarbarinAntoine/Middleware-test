package controllers

import "Middleware-test/internal/middlewares"

var IndexHandlerGetBundle = middlewares.Join(indexHandlerGet, middlewares.Log())
var IndexHandlerPostBundle = middlewares.Join(indexHandlerPost, middlewares.Log())
var IndexHandlerPutBundle = middlewares.Join(indexHandlerPut, middlewares.Log(), middlewares.Guard())
var IndexHandlerDeleteBundle = middlewares.Join(indexHandlerDelete, middlewares.Log(), middlewares.Guard(), middlewares.Foo())
var IndexHandlerNoMethBundle = middlewares.Join(indexHandlerNoMeth, middlewares.Log(), middlewares.Foo())
var IndexHandlerOtherBundle = middlewares.Join(indexHandlerOther, middlewares.Log(), middlewares.Foo())

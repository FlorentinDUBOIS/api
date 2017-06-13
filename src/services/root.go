package services

import (
	"github.com/FlorentinDUBOIS/bouncer/src/provider/repositories"
	"github.com/FlorentinDUBOIS/bouncer/src/services/factories"
)

var userRepository = new(repositories.UserRepository)
var apiUserFactory = new(factories.APIUserFactory)
var pgUserFactory = new(factories.PostgresUserFactory)

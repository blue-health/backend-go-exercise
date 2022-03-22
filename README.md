# Golang Backend Exercise

Welcome to the programming exercise for Golang Backend Developers! :)

## Your Task

Your task is to complete the implementation of this very simple Golang microservice. This microservice manages *coupon campaigns*, and allows the customer to apply a coupon to his *shopping cart*. A campaign is a collection of coupons. It also defines bolean *conditions* that check if a coupon can be applied to the given shopping cart, as well as *effects* which modify the shopping cart to apply the discount.

In particular, you should:

1. Finish the implementation of the business logic in the `loyalty` package
2. Finish the implementation of the REST API in the `web` package
3. Provide some unit tests for the above packages.

The implementation does not have to be production ready and fully tested, however it should show your reasoning and engineering practices.

## Setup

Please make sure you have a recent version of `golang` and `make` installed on your local machine.

You can run this service with `make run`, build it with `make build`, test it with `make test`. Please take a look into the `Makefile` to see all available operations.

## Contact

In case of any questions or problems you may have, please reach out to me at [marcin.praski@immunkarte.de](mailto:marcin.praski@immunkarte.de).

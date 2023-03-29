# Changelog

## [2.1.1](https://github.com/shipperizer/miniature-monkey/compare/v2.1.0...v2.1.1) (2023-03-29)


### Bug Fixes

* use plain request path instead of chi context RoutePath ([632226b](https://github.com/shipperizer/miniature-monkey/commit/632226b477967f5d4f6e81fa725f082c5dffe650))

## [2.1.0](https://github.com/shipperizer/miniature-monkey/compare/v2.0.0...v2.1.0) (2023-03-28)


### Features

* enhance status endpoint with build info ([cb2f726](https://github.com/shipperizer/miniature-monkey/commit/cb2f726d3393b68a3d6c4657ffd69f62de820130))
* inject tracer object ([7e96169](https://github.com/shipperizer/miniature-monkey/commit/7e961692053416bd773461e4c8d7c3066aff6633))
* introduce otel monitoring ([0795c2a](https://github.com/shipperizer/miniature-monkey/commit/0795c2a7ca2e0ab80dd25cad1b5871691c794974))

## [2.0.0](https://github.com/shipperizer/miniature-monkey/compare/v1.0.1...v2.0.0) (2023-03-22)


### ⚠ BREAKING CHANGES

* move to use go-chi

### Features

* add logging package ([356bb10](https://github.com/shipperizer/miniature-monkey/commit/356bb10eca7af219141f8c7d2dfaeb4017baff6e))
* add monitoring package ([2b63d3b](https://github.com/shipperizer/miniature-monkey/commit/2b63d3bd0efba61e545c23d5898f018135767ec0))
* move to use go-chi ([d380c60](https://github.com/shipperizer/miniature-monkey/commit/d380c604087553c36abcb55cf454f7bc05ffc28d))
* split middlewares into separate package ([ba19558](https://github.com/shipperizer/miniature-monkey/commit/ba19558a2633f9e839e349d9b8e9e1c0d9bbe055))


### Bug Fixes

* adjust to chi types ([ef3a4e0](https://github.com/shipperizer/miniature-monkey/commit/ef3a4e0f827bb9e1c6a6c42a945b73e338e436be))

### [1.0.1](https://www.github.com/shipperizer/miniature-monkey/compare/v1.0.0...v1.0.1) (2021-10-01)


### Bug Fixes

* adjust comments/docs ([fbc855b](https://www.github.com/shipperizer/miniature-monkey/commit/fbc855bffe4a523c9c82b88e5119085476d19864))

## 1.0.0 (2021-09-30)


### Features

* core module, containing principal functionality of library ([5a17244](https://www.github.com/shipperizer/miniature-monkey/commit/5a172442fcdab732a43df8e30badf052b794792d))
* webiface module, containing main interfaces of library ([ae8f655](https://www.github.com/shipperizer/miniature-monkey/commit/ae8f65500017fb59945ec9bc80bf79420abf16d0))


### Bug Fixes

* add extra modules for logging and return types ([9cd64d7](https://www.github.com/shipperizer/miniature-monkey/commit/9cd64d7152310e7ae9e21cf58c0c6f7a4193e22f))
* config module, holding APIConfig ([0df41b8](https://www.github.com/shipperizer/miniature-monkey/commit/0df41b8ecde415e3dbcb769893a7ae0025211548))
* create initial go.mod file with basic requirements ([5515c35](https://www.github.com/shipperizer/miniature-monkey/commit/5515c35d7225653179cb045c4caa783ae05621db))
* monitoring module, containing basic middleware funcs and basic prometheus client ([611c28f](https://www.github.com/shipperizer/miniature-monkey/commit/611c28f6b14b5c9df0c99c92a76020d658fe9481))
* status module, containing basic status endpoint ([df538e6](https://www.github.com/shipperizer/miniature-monkey/commit/df538e66b8e685f61a716b7ac494cf501c6b64bf))

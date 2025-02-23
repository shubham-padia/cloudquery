# Changelog

## [9.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v9.2.0...plugins-source-github-v9.2.1) (2024-04-09)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.38.0 ([#17554](https://github.com/cloudquery/cloudquery/issues/17554)) ([edb6f06](https://github.com/cloudquery/cloudquery/commit/edb6f066c3a3675f5bfca3e492eba3aeb31e770b))

## [9.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v9.1.0...plugins-source-github-v9.2.0) (2024-04-08)


### Features

* Add table option to filter workflow runs ([#17539](https://github.com/cloudquery/cloudquery/issues/17539)) ([202fe42](https://github.com/cloudquery/cloudquery/commit/202fe42c9a787c2deb4008fb8dd34752901637d9))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.4 ([#17485](https://github.com/cloudquery/cloudquery/issues/17485)) ([f370de4](https://github.com/cloudquery/cloudquery/commit/f370de449e61244398e6f413b973cbfa15c019a6))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.5 ([#17526](https://github.com/cloudquery/cloudquery/issues/17526)) ([554c499](https://github.com/cloudquery/cloudquery/commit/554c499eb9bc9f98f6f3dc4be23fd02049f48dcd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.37.0 ([#17532](https://github.com/cloudquery/cloudquery/issues/17532)) ([8080970](https://github.com/cloudquery/cloudquery/commit/8080970f40d22b6bc9db4c359780c744b476bb02))

## [9.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v9.0.0...plugins-source-github-v9.1.0) (2024-04-02)


### Features

* Add support for conditional requests via local cache ([#17368](https://github.com/cloudquery/cloudquery/issues/17368)) ([6d866d1](https://github.com/cloudquery/cloudquery/commit/6d866d18f3e8072eb83a633d0ae09aa852798f1a))


### Bug Fixes

* GitHub pagination handling ([#17462](https://github.com/cloudquery/cloudquery/issues/17462)) ([822684f](https://github.com/cloudquery/cloudquery/commit/822684f01aec6e72e579874243e8df2012661552))

## [9.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v8.2.1...plugins-source-github-v9.0.0) (2024-04-02)


### ⚠ BREAKING CHANGES

* Change default concurrency to 1500 instead of 10000 to better align with GitHub rate limits ([#17369](https://github.com/cloudquery/cloudquery/issues/17369))
* Skip archived repositories by default, rename `skip_archived_repos` to `include_archived_repos` ([#17339](https://github.com/cloudquery/cloudquery/issues/17339))

### Features

* Skip archived repositories by default, rename `skip_archived_repos` to `include_archived_repos` ([#17339](https://github.com/cloudquery/cloudquery/issues/17339)) ([b1ab433](https://github.com/cloudquery/cloudquery/commit/b1ab433925f5620b8f800b32678cb72f333bb8b2))


### Bug Fixes

* Change default concurrency to 1500 instead of 10000 to better align with GitHub rate limits ([#17369](https://github.com/cloudquery/cloudquery/issues/17369)) ([b757cf7](https://github.com/cloudquery/cloudquery/commit/b757cf72fe7c82e2cca599865b3b3d62e0330a7c))
* **deps:** Update module github.com/cloudquery/codegen to v0.3.13 ([#17444](https://github.com/cloudquery/cloudquery/issues/17444)) ([da276fe](https://github.com/cloudquery/cloudquery/commit/da276fe64c46ec0a5f182c83ebc32a90d55f5d50))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.1 ([#17372](https://github.com/cloudquery/cloudquery/issues/17372)) ([aaf6187](https://github.com/cloudquery/cloudquery/commit/aaf61873ae5d2e01ea5f3b8b319e4f79afb7b29c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.2 ([#17450](https://github.com/cloudquery/cloudquery/issues/17450)) ([2947506](https://github.com/cloudquery/cloudquery/commit/294750650269f8191c6dfff060c4d3a546405763))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.3 ([#17456](https://github.com/cloudquery/cloudquery/issues/17456)) ([020865a](https://github.com/cloudquery/cloudquery/commit/020865a6fde8c896947a844021f0cd7daeb01b06))

## [8.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v8.2.0...plugins-source-github-v8.2.1) (2024-03-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.0 ([#17325](https://github.com/cloudquery/cloudquery/issues/17325)) ([eb1b4b8](https://github.com/cloudquery/cloudquery/commit/eb1b4b8b963917b87ff644318cfec9745471d50a))

## [8.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v8.1.3...plugins-source-github-v8.2.0) (2024-03-22)


### Features

* When SBOM API returns 404 recommend enabling dependency graph ([#17309](https://github.com/cloudquery/cloudquery/issues/17309)) ([2c95674](https://github.com/cloudquery/cloudquery/commit/2c95674c516b1caee2d0e269afec2c65b0fc2998))


### Bug Fixes

* **deps:** Update github.com/cloudquery/jsonschema digest to 92878fa ([#16718](https://github.com/cloudquery/cloudquery/issues/16718)) ([7fe8588](https://github.com/cloudquery/cloudquery/commit/7fe858818fe1f88fcca6304c873a4614767a57b9))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.35.0 ([#17299](https://github.com/cloudquery/cloudquery/issues/17299)) ([524ba20](https://github.com/cloudquery/cloudquery/commit/524ba202801c2ae1eb59a5b462a5efc62d1b4000))

## [8.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v8.1.2...plugins-source-github-v8.1.3) (2024-03-19)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.0 ([#17203](https://github.com/cloudquery/cloudquery/issues/17203)) ([4b128b6](https://github.com/cloudquery/cloudquery/commit/4b128b6722dea883d66458f2f3c831184926353d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.1 ([#17220](https://github.com/cloudquery/cloudquery/issues/17220)) ([08d4950](https://github.com/cloudquery/cloudquery/commit/08d49504aee10f6883e1bd4f7e1102a274c8ee81))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.2 ([#17229](https://github.com/cloudquery/cloudquery/issues/17229)) ([41ed721](https://github.com/cloudquery/cloudquery/commit/41ed721cfa435a4937f3022501dd4d45a3a880b0))
* **deps:** Update module google.golang.org/protobuf to v1.33.0 [SECURITY] ([#17156](https://github.com/cloudquery/cloudquery/issues/17156)) ([a1e7465](https://github.com/cloudquery/cloudquery/commit/a1e7465f8df1c3423e323af9c0bff9d8e80638fe))
* Move GitHub `client.New` to Init function ([#17111](https://github.com/cloudquery/cloudquery/issues/17111)) ([9753f03](https://github.com/cloudquery/cloudquery/commit/9753f037a925b2db2fe3b2289ac37fb2819245bf))

## [8.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v8.1.1...plugins-source-github-v8.1.2) (2024-03-12)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.32.1 ([#17044](https://github.com/cloudquery/cloudquery/issues/17044)) ([d3592e7](https://github.com/cloudquery/cloudquery/commit/d3592e7f3ae600655778eb508aeccfa4e5b74e8c))

## [8.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v8.1.0...plugins-source-github-v8.1.1) (2024-03-05)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 814bf88 ([#16977](https://github.com/cloudquery/cloudquery/issues/16977)) ([d4d0e81](https://github.com/cloudquery/cloudquery/commit/d4d0e8138ec10e2c27eb0bf83e88905e838570d0))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to df926f6 ([#16980](https://github.com/cloudquery/cloudquery/issues/16980)) ([4684a2b](https://github.com/cloudquery/cloudquery/commit/4684a2b84b9c0f3c9dfd214b2cda517a46e8a0fb))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to df926f6 ([#16981](https://github.com/cloudquery/cloudquery/issues/16981)) ([4d6cef9](https://github.com/cloudquery/cloudquery/commit/4d6cef9134401b9a6fcd60e70683f1992e526c4d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.32.0 ([#16957](https://github.com/cloudquery/cloudquery/issues/16957)) ([8ffe2fe](https://github.com/cloudquery/cloudquery/commit/8ffe2fe13a11144cc4f463b01ede1d59c49fcc96))

## [8.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v8.0.2...plugins-source-github-v8.1.0) (2024-02-28)


### This Release has the Following Changes to Tables
- Table `github_workflows`: column added with name `contents_as_json` and type `json`

### Features

* Add `contents_as_json` column to `github_workflows` table ([#16846](https://github.com/cloudquery/cloudquery/issues/16846)) ([16d9db0](https://github.com/cloudquery/cloudquery/commit/16d9db0cd73fbd8994b59d86cdd0919e120db2e4))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.31.0 ([#16899](https://github.com/cloudquery/cloudquery/issues/16899)) ([2fac27a](https://github.com/cloudquery/cloudquery/commit/2fac27a2e3e789f6152b643c0af1c97ee95c4745))

## [8.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v8.0.1...plugins-source-github-v8.0.2) (2024-02-26)


### Bug Fixes

* GitHub app authentication with GitHub enterprise ([#16872](https://github.com/cloudquery/cloudquery/issues/16872)) ([0f3becf](https://github.com/cloudquery/cloudquery/commit/0f3becf4d122559f73db73d363becc661a35604b))
* Update GitHub example config ([#16841](https://github.com/cloudquery/cloudquery/issues/16841)) ([fb65286](https://github.com/cloudquery/cloudquery/commit/fb65286719f2924f3be7fcc8d45b01a95fa09c97))

## [8.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v8.0.0...plugins-source-github-v8.0.1) (2024-02-23)


### Bug Fixes

* Pass enterprise base URL to app auth if needed ([#16827](https://github.com/cloudquery/cloudquery/issues/16827)) ([bb953c9](https://github.com/cloudquery/cloudquery/commit/bb953c9146351147648bfa19106bfa1772a24c1c))

## [8.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.7.1...plugins-source-github-v8.0.0) (2024-02-21)


### This Release has the Following Changes to Tables
- Table `github_billing_action`: column type changed from `int64` to `float64` for `included_minutes` (:warning: breaking)
- Table `github_billing_action`: column type changed from `int64` to `float64` for `total_minutes_used` (:warning: breaking)
- Table `github_billing_package`: column type changed from `int64` to `float64` for `included_gigabytes_bandwidth` (:warning: breaking)
- Table `github_billing_storage`: column type changed from `int64` to `float64` for `estimated_storage_for_month` (:warning: breaking)
- Table `github_issues`: column added with name `draft` and type `bool`
- Table `github_organization_dependabot_alerts`: column added with name `auto_dismissed_at` and type `timestamp[us, tz=UTC]`
- Table `github_organization_dependabot_alerts`: column added with name `repository` and type `json`
- Table `github_organizations`: column added with name `secret_scanning_validity_checks_enabled` and type `bool`
- Table `github_repositories`: column added with name `custom_properties` and type `json`
- Table `github_repositories`: column added with name `web_commit_signoff_required` and type `bool`
- Table `github_repository_dependabot_alerts`: column added with name `auto_dismissed_at` and type `timestamp[us, tz=UTC]`
- Table `github_repository_dependabot_alerts`: column added with name `repository` and type `json`
- Table `github_repository_keys`: column added with name `added_by` and type `utf8`
- Table `github_repository_keys`: column added with name `last_used` and type `timestamp[us, tz=UTC]`
- Table `github_repository_sboms` was added
- Table `github_team_repositories`: column added with name `custom_properties` and type `json`
- Table `github_team_repositories`: column added with name `web_commit_signoff_required` and type `bool`
- Table `github_workflow_jobs`: column added with name `created_at` and type `timestamp[us, tz=UTC]`
- Table `github_workflow_jobs`: column added with name `head_branch` and type `utf8`
- Table `github_workflow_jobs`: column added with name `workflow_name` and type `utf8`
- Table `github_workflow_runs`: column added with name `display_title` and type `utf8`
- Table `github_workflow_runs`: column added with name `referenced_workflows` and type `json`
- Table `github_workflow_runs`: column added with name `triggering_actor` and type `json`

### ⚠ BREAKING CHANGES

* Bump `go-github` to v59 in Github plugin ([#16797](https://github.com/cloudquery/cloudquery/issues/16797))

### Features

* Add `github_repository_sboms` table to sync Github SBOM data ([#16796](https://github.com/cloudquery/cloudquery/issues/16796)) ([992434f](https://github.com/cloudquery/cloudquery/commit/992434f7e3f2c90d4f84d53495907f298f3cee31))
* Bump `go-github` to v59 in Github plugin ([#16797](https://github.com/cloudquery/cloudquery/issues/16797)) ([5b89ef6](https://github.com/cloudquery/cloudquery/commit/5b89ef641d415c6f88dce1a461ab05f7c1625ff3))

## [7.7.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.7.0...plugins-source-github-v7.7.1) (2024-02-20)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.30.0 ([#16669](https://github.com/cloudquery/cloudquery/issues/16669)) ([44b9729](https://github.com/cloudquery/cloudquery/commit/44b9729fa5d7590f65b9073ce4a1fc18a529117e))

## [7.7.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.6.4...plugins-source-github-v7.7.0) (2024-02-14)


### Features

* Adding JSONSchema support to GitHub plugin ([#16491](https://github.com/cloudquery/cloudquery/issues/16491)) ([d81cb01](https://github.com/cloudquery/cloudquery/commit/d81cb01a8ebed190bde1ec641841b3e57d71d574))


### Bug Fixes

* **deps:** Update github.com/beatlabs/github-auth digest to 98862c3 ([#16416](https://github.com/cloudquery/cloudquery/issues/16416)) ([1d3b946](https://github.com/cloudquery/cloudquery/commit/1d3b9467291a4cbac996c9ae565cb26ce4d435a6))
* **deps:** Update github.com/cloudquery/jsonschema digest to d771afd ([#16527](https://github.com/cloudquery/cloudquery/issues/16527)) ([7c26840](https://github.com/cloudquery/cloudquery/commit/7c2684039a6d774af466c581d93aae4cda00f412))
* **deps:** Update golang.org/x/exp digest to 1b97071 ([#16419](https://github.com/cloudquery/cloudquery/issues/16419)) ([6d77cd1](https://github.com/cloudquery/cloudquery/commit/6d77cd19b6fc648a4ddb12025c22127e960036a4))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 1f4bbc5 ([#16421](https://github.com/cloudquery/cloudquery/issues/16421)) ([9489931](https://github.com/cloudquery/cloudquery/commit/9489931c1b64bf1f7d5da51997944ee54370215b))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to 1f4bbc5 ([#16422](https://github.com/cloudquery/cloudquery/issues/16422)) ([74e98fc](https://github.com/cloudquery/cloudquery/commit/74e98fcbde6c6e11baf98284aef0341c597d4817))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.0 ([#16237](https://github.com/cloudquery/cloudquery/issues/16237)) ([3fcdab0](https://github.com/cloudquery/cloudquery/commit/3fcdab08816ad9de7bb4eecab59c7be1bda3d00c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.1 ([#16296](https://github.com/cloudquery/cloudquery/issues/16296)) ([ab4a0da](https://github.com/cloudquery/cloudquery/commit/ab4a0dace0a870755fd22d92c6e9c999351f594e))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.2 ([#16342](https://github.com/cloudquery/cloudquery/issues/16342)) ([f3eb857](https://github.com/cloudquery/cloudquery/commit/f3eb85729e5db16c2530b31d6d276934866d5ef0))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.28.0 ([#16362](https://github.com/cloudquery/cloudquery/issues/16362)) ([9166b6b](https://github.com/cloudquery/cloudquery/commit/9166b6b603d0d56a30c2e5072c4f2da5c0c837b5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.29.0 ([#16395](https://github.com/cloudquery/cloudquery/issues/16395)) ([fb1102e](https://github.com/cloudquery/cloudquery/commit/fb1102eac8af4b3722b82b882187fdf322546513))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.29.1 ([#16430](https://github.com/cloudquery/cloudquery/issues/16430)) ([738e89f](https://github.com/cloudquery/cloudquery/commit/738e89f2c969a8a3f1698a8686aeaddb358e7a23))

## [7.6.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.6.3...plugins-source-github-v7.6.4) (2024-01-16)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v15 digest to 6d44906 ([#16115](https://github.com/cloudquery/cloudquery/issues/16115)) ([8b0ae62](https://github.com/cloudquery/cloudquery/commit/8b0ae6266d19a10fe84102837802358f0b9bb1bc))
* **deps:** Update github.com/apache/arrow/go/v15 digest to 7e703aa ([#16134](https://github.com/cloudquery/cloudquery/issues/16134)) ([72d5eb3](https://github.com/cloudquery/cloudquery/commit/72d5eb35644ce78d775790b0298a0c7690788d28))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.1 ([#16069](https://github.com/cloudquery/cloudquery/issues/16069)) ([edda65c](https://github.com/cloudquery/cloudquery/commit/edda65c238b2cb78a7a2078b62557a7d8d822e49))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.2 ([#16130](https://github.com/cloudquery/cloudquery/issues/16130)) ([7ae6f41](https://github.com/cloudquery/cloudquery/commit/7ae6f41957edb3446ff3175857aaf3dcea2cf5bc))

## [7.6.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.6.2...plugins-source-github-v7.6.3) (2024-01-04)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.0 ([#15932](https://github.com/cloudquery/cloudquery/issues/15932)) ([2292b5a](https://github.com/cloudquery/cloudquery/commit/2292b5a2aa5936f2529238a05708de0b3bde9a35))

## [7.6.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.6.1...plugins-source-github-v7.6.2) (2024-01-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v15 digest to 7c3480e ([#15904](https://github.com/cloudquery/cloudquery/issues/15904)) ([a3ec012](https://github.com/cloudquery/cloudquery/commit/a3ec01203183e5c94630beae86434519e87e225d))
* **deps:** Update github.com/beatlabs/github-auth digest to 3b7665f ([#15906](https://github.com/cloudquery/cloudquery/issues/15906)) ([5fe10da](https://github.com/cloudquery/cloudquery/commit/5fe10dade4fa692f329cfb6355b3b73e70027ba0))
* **deps:** Update github.com/gomarkdown/markdown digest to 1d6d208 ([#15907](https://github.com/cloudquery/cloudquery/issues/15907)) ([86d29a9](https://github.com/cloudquery/cloudquery/commit/86d29a900e6c9dbcad09f5b0c4b0615aee59a2ae))
* **deps:** Update golang.org/x/exp digest to 02704c9 ([#15909](https://github.com/cloudquery/cloudquery/issues/15909)) ([dfe32d2](https://github.com/cloudquery/cloudquery/commit/dfe32d2557dcac0fb6dc741c9df4edccdcb07076))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 995d672 ([#15911](https://github.com/cloudquery/cloudquery/issues/15911)) ([18ac2b8](https://github.com/cloudquery/cloudquery/commit/18ac2b806d798e0a9052cc10e8442557ab1c4253))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.24.0 ([#15863](https://github.com/cloudquery/cloudquery/issues/15863)) ([47d7899](https://github.com/cloudquery/cloudquery/commit/47d78994370f083912b6d4329f12d5cef9c255d5))

## [7.6.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.6.0...plugins-source-github-v7.6.1) (2023-12-28)


### Bug Fixes

* **deps:** Update `github.com/apache/arrow/go` to `v15` ([#15754](https://github.com/cloudquery/cloudquery/issues/15754)) ([bd962eb](https://github.com/cloudquery/cloudquery/commit/bd962eb1093cf09e928e2a0e7782288ec4020ec4))
* **deps:** Update github.com/apache/arrow/go/v15 digest to bcaeaa8 ([#15791](https://github.com/cloudquery/cloudquery/issues/15791)) ([89dc812](https://github.com/cloudquery/cloudquery/commit/89dc81201529de2a1fc1ecce5efa74d6f363e57b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.0 ([#15651](https://github.com/cloudquery/cloudquery/issues/15651)) ([6e96125](https://github.com/cloudquery/cloudquery/commit/6e96125a9d9c75616483952edb7a9e402818b264))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.1 ([#15699](https://github.com/cloudquery/cloudquery/issues/15699)) ([67c10c3](https://github.com/cloudquery/cloudquery/commit/67c10c38a04dcdd1512bf6dc739f89bc11baa888))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.2 ([#15724](https://github.com/cloudquery/cloudquery/issues/15724)) ([ad750b1](https://github.com/cloudquery/cloudquery/commit/ad750b1530af06353f2225c7d3397af580093687))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.23.0 ([#15803](https://github.com/cloudquery/cloudquery/issues/15803)) ([b6f9373](https://github.com/cloudquery/cloudquery/commit/b6f937385020c63ce59b2bc60402752b6c239c6c))
* **deps:** Update module golang.org/x/crypto to v0.17.0 [SECURITY] ([#15730](https://github.com/cloudquery/cloudquery/issues/15730)) ([718be50](https://github.com/cloudquery/cloudquery/commit/718be502014ff36aa50cde3a83453b3d6ce15a83))

## [7.6.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.5.1...plugins-source-github-v7.6.0) (2023-12-11)


### Features

* Add option to skip archived repositories ([#15593](https://github.com/cloudquery/cloudquery/issues/15593)) ([9e69ad8](https://github.com/cloudquery/cloudquery/commit/9e69ad8248b81044d1fa944297d257fd1dfe13c1))


### Bug Fixes

* **deps:** Update github.com/beatlabs/github-auth digest to b1420bd ([#15514](https://github.com/cloudquery/cloudquery/issues/15514)) ([40d2882](https://github.com/cloudquery/cloudquery/commit/40d28823beae51372fdd2c99cefcf77da7333c8c))
* **deps:** Update github.com/gomarkdown/markdown digest to a660076 ([#15517](https://github.com/cloudquery/cloudquery/issues/15517)) ([fa1334c](https://github.com/cloudquery/cloudquery/commit/fa1334c5ce0e157834b0cd676b38af26510fbe43))
* **deps:** Update golang.org/x/exp digest to 6522937 ([#15518](https://github.com/cloudquery/cloudquery/issues/15518)) ([69f9a06](https://github.com/cloudquery/cloudquery/commit/69f9a06754b2feb7c73bd5a80d42fd191c7fdb17))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 3a041ad ([#15520](https://github.com/cloudquery/cloudquery/issues/15520)) ([b2a322a](https://github.com/cloudquery/cloudquery/commit/b2a322a5ec5c1945af5a655c759493a879a9be09))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.2 ([#15325](https://github.com/cloudquery/cloudquery/issues/15325)) ([77f2db5](https://github.com/cloudquery/cloudquery/commit/77f2db52634bad6e56d970d55172b08d823b97c9))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.3 ([#15355](https://github.com/cloudquery/cloudquery/issues/15355)) ([d8455e5](https://github.com/cloudquery/cloudquery/commit/d8455e5ca1059746c7aced395e9bc150ea495591))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.19.0 ([#15509](https://github.com/cloudquery/cloudquery/issues/15509)) ([41c689d](https://github.com/cloudquery/cloudquery/commit/41c689d0835487a8d924bb11c989c231f5e3df7c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.19.1 ([#15539](https://github.com/cloudquery/cloudquery/issues/15539)) ([a298555](https://github.com/cloudquery/cloudquery/commit/a298555343fc7ad483361c2f19c3d39693dab882))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.20.0 ([#15574](https://github.com/cloudquery/cloudquery/issues/15574)) ([317dca4](https://github.com/cloudquery/cloudquery/commit/317dca4182478d6f3789082ae563d9e8bd417d20))

## [7.5.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.5.0...plugins-source-github-v7.5.1) (2023-11-16)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.1 ([#15263](https://github.com/cloudquery/cloudquery/issues/15263)) ([a9a39ef](https://github.com/cloudquery/cloudquery/commit/a9a39efe8112a564f21c06ba7627fe6c7ced4cdf))

## [7.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.4.2...plugins-source-github-v7.5.0) (2023-11-15)


### Features

* Add more details to GitHub source app auth ([#15134](https://github.com/cloudquery/cloudquery/issues/15134)) ([f56257c](https://github.com/cloudquery/cloudquery/commit/f56257ca3c6678ce5ef255723f0d3af7a092a62e))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.4.2 ([#15034](https://github.com/cloudquery/cloudquery/issues/15034)) ([45c2caa](https://github.com/cloudquery/cloudquery/commit/45c2caa345aa33199ad1592bf378a5a839612c6f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.17.0 ([#15064](https://github.com/cloudquery/cloudquery/issues/15064)) ([9c2db8c](https://github.com/cloudquery/cloudquery/commit/9c2db8cedaec682a89b444db29e8c0fb45989408))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.17.1 ([#15075](https://github.com/cloudquery/cloudquery/issues/15075)) ([151769e](https://github.com/cloudquery/cloudquery/commit/151769e7c02028a04ef0ed280951c000ebb1f9c2))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.0 ([#15132](https://github.com/cloudquery/cloudquery/issues/15132)) ([81ee138](https://github.com/cloudquery/cloudquery/commit/81ee138ff86c4b92c3ec93208e0a7e05af2b0036))
* **deps:** Update module github.com/docker/docker to v24 [SECURITY] ([#15060](https://github.com/cloudquery/cloudquery/issues/15060)) ([41acd0e](https://github.com/cloudquery/cloudquery/commit/41acd0e4ac63221e90cca89a7137a8685692267d))
* **deps:** Update module google.golang.org/grpc to v1.58.3 [SECURITY] ([#14940](https://github.com/cloudquery/cloudquery/issues/14940)) ([e1addea](https://github.com/cloudquery/cloudquery/commit/e1addeaf58ad965e545a3e068860609dadcffa10))

## [7.4.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.4.1...plugins-source-github-v7.4.2) (2023-10-24)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to f46436f ([#14803](https://github.com/cloudquery/cloudquery/issues/14803)) ([f5248d7](https://github.com/cloudquery/cloudquery/commit/f5248d749398ded6a50903e09ecabbb996e94a34))

## [7.4.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.4.0...plugins-source-github-v7.4.1) (2023-10-19)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.9 ([#14627](https://github.com/cloudquery/cloudquery/issues/14627)) ([c1d244c](https://github.com/cloudquery/cloudquery/commit/c1d244c95199141ac39a713a3f0577b2fb3bf736))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.3.0 ([#14635](https://github.com/cloudquery/cloudquery/issues/14635)) ([00b380c](https://github.com/cloudquery/cloudquery/commit/00b380c10be1642f737f871ba5588888ed5dd180))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.4.0 ([#14639](https://github.com/cloudquery/cloudquery/issues/14639)) ([f139c0e](https://github.com/cloudquery/cloudquery/commit/f139c0e9369ef92a3cd874003db40b48e229ab58))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.14.0 ([#14577](https://github.com/cloudquery/cloudquery/issues/14577)) ([223c4c1](https://github.com/cloudquery/cloudquery/commit/223c4c1df6c432d7f1bf67a48114e417282bcd0f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.0 ([#14622](https://github.com/cloudquery/cloudquery/issues/14622)) ([b497a6b](https://github.com/cloudquery/cloudquery/commit/b497a6bc5645854bd25d4083fd91ec549a7f274f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.2 ([#14662](https://github.com/cloudquery/cloudquery/issues/14662)) ([e274fe4](https://github.com/cloudquery/cloudquery/commit/e274fe419f6cacdf62547cd7134f40916e5ddd96))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.3 ([#14679](https://github.com/cloudquery/cloudquery/issues/14679)) ([0513c19](https://github.com/cloudquery/cloudquery/commit/0513c193919f4555d41f22ba2ff66efaaf5fca67))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.16.1 ([#14721](https://github.com/cloudquery/cloudquery/issues/14721)) ([1c7ee1d](https://github.com/cloudquery/cloudquery/commit/1c7ee1dc99d7a9cb3358a83e8d827d59be78cefa))
* Set plugin metadata ([#14715](https://github.com/cloudquery/cloudquery/issues/14715)) ([39935e2](https://github.com/cloudquery/cloudquery/commit/39935e2531c4edbd960d5db91e1027b13d7c0a4f))
* Update plugin-SDK to v4.16.0 ([#14702](https://github.com/cloudquery/cloudquery/issues/14702)) ([0dcb545](https://github.com/cloudquery/cloudquery/commit/0dcb5455a71eaa7d28193b1b2fbcdd184dfad2ab))

## [7.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.3.4...plugins-source-github-v7.4.0) (2023-10-16)


### This Release has the Following Changes to Tables
- Table `github_workflow_jobs` was added
- Table `github_workflow_run_usage` was added
- Table `github_workflow_runs` was added

### Features

* Adding GitHub Actions workflowRuns, RunUsage, and Jobs ([#14529](https://github.com/cloudquery/cloudquery/issues/14529)) ([9e91609](https://github.com/cloudquery/cloudquery/commit/9e916090240a33ddf5a0163b950bf8fc1a8873ed))


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to d401686 ([#14459](https://github.com/cloudquery/cloudquery/issues/14459)) ([7ce40f8](https://github.com/cloudquery/cloudquery/commit/7ce40f8dcb1e408c385e877e56b5bb78906b10d2))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to dbcb149 ([#14537](https://github.com/cloudquery/cloudquery/issues/14537)) ([68686f4](https://github.com/cloudquery/cloudquery/commit/68686f4e7636db02bddd961e3d75b60d5218ca85))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.6 ([#14475](https://github.com/cloudquery/cloudquery/issues/14475)) ([83fe7ca](https://github.com/cloudquery/cloudquery/commit/83fe7ca2f5fa83bd3219ddde8fe44fcf1d447480))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.8 ([#14503](https://github.com/cloudquery/cloudquery/issues/14503)) ([4056593](https://github.com/cloudquery/cloudquery/commit/40565937cfc12b33048980b55e91a9a60a62bd47))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.2 ([#14378](https://github.com/cloudquery/cloudquery/issues/14378)) ([a2e0c46](https://github.com/cloudquery/cloudquery/commit/a2e0c4615af4aa205fa082d3f196ea2dc5ce2445))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.3 ([#14436](https://github.com/cloudquery/cloudquery/issues/14436)) ([d529e2d](https://github.com/cloudquery/cloudquery/commit/d529e2d22da93a234492c4165e7eed1257c5767f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.4 ([#14489](https://github.com/cloudquery/cloudquery/issues/14489)) ([9bb45dc](https://github.com/cloudquery/cloudquery/commit/9bb45dc2dacc2c7a6fbd47538b954f731741809b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.5 ([#14516](https://github.com/cloudquery/cloudquery/issues/14516)) ([2d905bf](https://github.com/cloudquery/cloudquery/commit/2d905bf9ea81556282c8ca27dcc6334606a2e83b))
* **deps:** Update module golang.org/x/net to v0.17.0 [SECURITY] ([#14500](https://github.com/cloudquery/cloudquery/issues/14500)) ([9e603d5](https://github.com/cloudquery/cloudquery/commit/9e603d50d28033ed5bf451e569abc7c25014dbfb))

## [7.3.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.3.3...plugins-source-github-v7.3.4) (2023-10-04)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v14 digest to 00efb06 ([#14202](https://github.com/cloudquery/cloudquery/issues/14202)) ([fc8cc62](https://github.com/cloudquery/cloudquery/commit/fc8cc62ed70db157612e88678c123ba6a34b3b3c))
* **deps:** Update github.com/beatlabs/github-auth digest to cdaa33a ([#14213](https://github.com/cloudquery/cloudquery/issues/14213)) ([9e53b36](https://github.com/cloudquery/cloudquery/commit/9e53b3683b86e4d8aedef692175c2ceb34cc9c51))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 7ded38b ([#14246](https://github.com/cloudquery/cloudquery/issues/14246)) ([005891e](https://github.com/cloudquery/cloudquery/commit/005891e1892b41235ddb3b102f4bb6dafd48949a))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.0 ([#14281](https://github.com/cloudquery/cloudquery/issues/14281)) ([85835a9](https://github.com/cloudquery/cloudquery/commit/85835a938bfa58d1b0d320ae17aff5fe7f6cfef2))

## [7.3.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.3.2...plugins-source-github-v7.3.3) (2023-09-28)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to ffb7089 ([#13879](https://github.com/cloudquery/cloudquery/issues/13879)) ([f95ced5](https://github.com/cloudquery/cloudquery/commit/f95ced5daa2b123bd71ddff75bd76b3b008790c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.0 ([#13978](https://github.com/cloudquery/cloudquery/issues/13978)) ([2efdf55](https://github.com/cloudquery/cloudquery/commit/2efdf55aed94a14c35c51632ff61ed454caaf5a5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.2 ([#13988](https://github.com/cloudquery/cloudquery/issues/13988)) ([aebaddf](https://github.com/cloudquery/cloudquery/commit/aebaddfc5ca0d7574b8cd72e9e074ec612472dbe))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.0 ([#14031](https://github.com/cloudquery/cloudquery/issues/14031)) ([ac7cdc4](https://github.com/cloudquery/cloudquery/commit/ac7cdc4f7d71599dad89b3170bb7bda676984228))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.1 ([#14063](https://github.com/cloudquery/cloudquery/issues/14063)) ([5a0ff7b](https://github.com/cloudquery/cloudquery/commit/5a0ff7b67890478c371385b379e0a8ef0c2f4865))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.8.0 ([#13950](https://github.com/cloudquery/cloudquery/issues/13950)) ([15b0b69](https://github.com/cloudquery/cloudquery/commit/15b0b6925932613ed2915a3255b3466f21a5c7bf))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.9.0 ([#13960](https://github.com/cloudquery/cloudquery/issues/13960)) ([f074076](https://github.com/cloudquery/cloudquery/commit/f074076a21dc0b8cadfdc3adb9731473d24d28b1))

## [7.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.3.1...plugins-source-github-v7.3.2) (2023-09-12)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 483f6b2 ([#13780](https://github.com/cloudquery/cloudquery/issues/13780)) ([8d31b44](https://github.com/cloudquery/cloudquery/commit/8d31b44f787f42d47f186cdcc4a5739a3a370a5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.0 ([#13625](https://github.com/cloudquery/cloudquery/issues/13625)) ([bb5463f](https://github.com/cloudquery/cloudquery/commit/bb5463fb5919f50f1327eecae884b2ab99fb8b34))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.1 ([#13713](https://github.com/cloudquery/cloudquery/issues/13713)) ([73004dc](https://github.com/cloudquery/cloudquery/commit/73004dcabd05bf474d8b5960b8c747a894b98560))

## [7.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.3.0...plugins-source-github-v7.3.1) (2023-09-05)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v14 digest to a526ba6 ([#13562](https://github.com/cloudquery/cloudquery/issues/13562)) ([248672b](https://github.com/cloudquery/cloudquery/commit/248672beb020828cde1cb608d5c1ed6d656c777b))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to cd3d411 ([#13598](https://github.com/cloudquery/cloudquery/issues/13598)) ([f22bfa6](https://github.com/cloudquery/cloudquery/commit/f22bfa6b2d4fd0caeacf0726ccd307db38f8860c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.4 ([#13519](https://github.com/cloudquery/cloudquery/issues/13519)) ([9d25165](https://github.com/cloudquery/cloudquery/commit/9d25165820703844c6de96688d939aa5033608ae))

## [7.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.2.1...plugins-source-github-v7.3.0) (2023-09-01)


### Features

* Support reading private key from string ([#13556](https://github.com/cloudquery/cloudquery/issues/13556)) ([9909e24](https://github.com/cloudquery/cloudquery/commit/9909e245a6e40dbcf5ea50505e06856367638c03))


### Bug Fixes

* **deps:** Update github.com/99designs/go-keychain digest to 9cf53c8 ([#13561](https://github.com/cloudquery/cloudquery/issues/13561)) ([a170256](https://github.com/cloudquery/cloudquery/commit/a17025657e92b017fe3c8bd37abfaa2354e6e818))
* **deps:** Update github.com/beatlabs/github-auth digest to 39adf75 ([#13563](https://github.com/cloudquery/cloudquery/issues/13563)) ([ea5f471](https://github.com/cloudquery/cloudquery/commit/ea5f4712aa07de2b25b7b9fe44c0c70a10674bc8))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.0 ([#13492](https://github.com/cloudquery/cloudquery/issues/13492)) ([c305876](https://github.com/cloudquery/cloudquery/commit/c305876e3d92944aa6c1a26547f786fdc5b50e23))

## [7.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.2.0...plugins-source-github-v7.2.1) (2023-08-29)


### Bug Fixes

* **deps:** Update `github.com/cloudquery/arrow/go/v13` to `github.com/apache/arrow/go/v14` ([#13341](https://github.com/cloudquery/cloudquery/issues/13341)) ([feb8f87](https://github.com/cloudquery/cloudquery/commit/feb8f87d8d761eb9c49ce84329ad0397f730a918))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 5b83d4f ([#13203](https://github.com/cloudquery/cloudquery/issues/13203)) ([b0a4b8c](https://github.com/cloudquery/cloudquery/commit/b0a4b8ccf7c429bf5a6ed88866865212015b68e4))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.1 ([#13195](https://github.com/cloudquery/cloudquery/issues/13195)) ([a184c37](https://github.com/cloudquery/cloudquery/commit/a184c3786ad49df8564344773e9b96f617ef87a1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.2 ([#13262](https://github.com/cloudquery/cloudquery/issues/13262)) ([5c55aa3](https://github.com/cloudquery/cloudquery/commit/5c55aa35282786375e8ce9493b2a4878e0fb27bc))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.5 ([#13285](https://github.com/cloudquery/cloudquery/issues/13285)) ([e076abd](https://github.com/cloudquery/cloudquery/commit/e076abd9d67813a29ced0c1b7b1664fd728b9ba8))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.6 ([#13345](https://github.com/cloudquery/cloudquery/issues/13345)) ([a995a05](https://github.com/cloudquery/cloudquery/commit/a995a0598a209e0fe3ba09f4ced2a052dc14b67a))

## [7.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.1.4...plugins-source-github-v7.2.0) (2023-08-18)


### Features

* Add `discovery_concurrency` sync option ([#13188](https://github.com/cloudquery/cloudquery/issues/13188)) ([a2380c1](https://github.com/cloudquery/cloudquery/commit/a2380c14ee7110831777e4c216d78af9acab069c))

## [7.1.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.1.3...plugins-source-github-v7.1.4) (2023-08-18)


### Bug Fixes

* Re-init `RepositoryListByOrgOptions` for each organization when listing repositories ([#13182](https://github.com/cloudquery/cloudquery/issues/13182)) ([cc08e74](https://github.com/cloudquery/cloudquery/commit/cc08e74fcf91063376f2bae1ec550164952c88e1))

## [7.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.1.2...plugins-source-github-v7.1.3) (2023-08-15)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e9683e1 ([#13015](https://github.com/cloudquery/cloudquery/issues/13015)) ([6557696](https://github.com/cloudquery/cloudquery/commit/65576966d3bd14297499a5b85d3b4fc2c7918df3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.0 ([#13068](https://github.com/cloudquery/cloudquery/issues/13068)) ([7bb0e4b](https://github.com/cloudquery/cloudquery/commit/7bb0e4ba654971726e16a6a501393e3831170307))

## [7.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.1.1...plugins-source-github-v7.1.2) (2023-08-10)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.4.0 ([#12850](https://github.com/cloudquery/cloudquery/issues/12850)) ([0861200](https://github.com/cloudquery/cloudquery/commit/086120054b45213947e95be954ba6164b9cf6587))

## [7.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.1.0...plugins-source-github-v7.1.1) (2023-08-08)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to 112f949 ([#12659](https://github.com/cloudquery/cloudquery/issues/12659)) ([48d73a9](https://github.com/cloudquery/cloudquery/commit/48d73a93e678994f43171c363f5a75c29547b0b9))
* **deps:** Update github.com/beatlabs/github-auth digest to 88fe74f ([#12663](https://github.com/cloudquery/cloudquery/issues/12663)) ([206912e](https://github.com/cloudquery/cloudquery/commit/206912e0c5f9bcfd0a11c62818793435c477a20b))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 3452eb0 ([#12595](https://github.com/cloudquery/cloudquery/issues/12595)) ([c1c0949](https://github.com/cloudquery/cloudquery/commit/c1c09490b17f2e64435e05d745890cdb8b22310d))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f53878d ([#12778](https://github.com/cloudquery/cloudquery/issues/12778)) ([6f5d58e](https://github.com/cloudquery/cloudquery/commit/6f5d58e3b84d3c76b1d1a3d6c5a488f77995a057))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.4 ([#12718](https://github.com/cloudquery/cloudquery/issues/12718)) ([f059a15](https://github.com/cloudquery/cloudquery/commit/f059a159a2ee406ab2b0a33792c244cd217025a6))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.5 ([#12731](https://github.com/cloudquery/cloudquery/issues/12731)) ([d267239](https://github.com/cloudquery/cloudquery/commit/d267239aa3aca5f94bd36a8db1ec0d9f7dc0865f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.6 ([#12799](https://github.com/cloudquery/cloudquery/issues/12799)) ([fb0e0d7](https://github.com/cloudquery/cloudquery/commit/fb0e0d75ab010f421c834e58d93676de76fcb423))

## [7.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.0.1...plugins-source-github-v7.1.0) (2023-07-25)


### This Release has the Following Changes to Tables
- Table `github_repository_keys` was added

### Features

* **resources:** Add Repository Keys ([#12407](https://github.com/cloudquery/cloudquery/issues/12407)) ([cce0cac](https://github.com/cloudquery/cloudquery/commit/cce0cacaf44aee203245a54846c2a2209177ceef))


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 10df4b9 ([#12443](https://github.com/cloudquery/cloudquery/issues/12443)) ([e385283](https://github.com/cloudquery/cloudquery/commit/e38528309f862f37bc7e278f9b69cf92d5aa5bd5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.3 ([#12307](https://github.com/cloudquery/cloudquery/issues/12307)) ([8f14e4d](https://github.com/cloudquery/cloudquery/commit/8f14e4de7bf4d4c833f501135ea0610916a42f8b))

## [7.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v7.0.0...plugins-source-github-v7.0.1) (2023-07-18)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8e2219b ([#12220](https://github.com/cloudquery/cloudquery/issues/12220)) ([24e8fb5](https://github.com/cloudquery/cloudquery/commit/24e8fb588740896fe11a660e8b80231e741b753c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.0 ([#12256](https://github.com/cloudquery/cloudquery/issues/12256)) ([eaec331](https://github.com/cloudquery/cloudquery/commit/eaec33165345ad51fdb6ddbffbf8a1199ebd6384))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.1 ([#12272](https://github.com/cloudquery/cloudquery/issues/12272)) ([557ca69](https://github.com/cloudquery/cloudquery/commit/557ca69a7dee9dabb80e6afb6f41f205fd8a80d8))

## [7.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v6.0.3...plugins-source-github-v7.0.0) (2023-07-14)


### ⚠ BREAKING CHANGES

* Upgrades the github source plugin to use plugin-sdk v4. This version does not contain any user-facing breaking changes, but because it is now using CloudQuery gRPC protocol v3, it does require use of a destination plugin that also supports protocol v3. All recent destination plugin versions support this.

### Features

* Upgrades the github source plugin to use plugin-sdk v4. This version does not contain any user-facing breaking changes, but because it is now using CloudQuery gRPC protocol v3, it does require use of a destination plugin that also supports protocol v3. All recent destination plugin versions support this. ([3ff5814](https://github.com/cloudquery/cloudquery/commit/3ff58141cbbe461ab515d4ff0f3d22892513bf08))


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0a52533 ([#12091](https://github.com/cloudquery/cloudquery/issues/12091)) ([927cefa](https://github.com/cloudquery/cloudquery/commit/927cefa943ec3969a2ec39b628bc1eba545a2108))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a2a76eb ([#12104](https://github.com/cloudquery/cloudquery/issues/12104)) ([311f474](https://github.com/cloudquery/cloudquery/commit/311f4749af2491a606f29483190717a5fe238da6))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.7.0 ([#12166](https://github.com/cloudquery/cloudquery/issues/12166)) ([94390dd](https://github.com/cloudquery/cloudquery/commit/94390dde19d0c37fee9d035219d62f6ae7edb127))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.1.0 ([#12174](https://github.com/cloudquery/cloudquery/issues/12174)) ([80f0289](https://github.com/cloudquery/cloudquery/commit/80f02892a4cf876c4bf4dd4fd9367afb3770ad26))
* **deps:** Upgrade source plugins to SDK v4.0.0 release ([#12135](https://github.com/cloudquery/cloudquery/issues/12135)) ([c20a111](https://github.com/cloudquery/cloudquery/commit/c20a111d591101fb1bbc42292accc953af38e8a6))

## [6.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v6.0.2...plugins-source-github-v6.0.3) (2023-07-04)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to 5a06b2e ([#11857](https://github.com/cloudquery/cloudquery/issues/11857)) ([43c2f5f](https://github.com/cloudquery/cloudquery/commit/43c2f5f3a893e5286f67c4943a9d1bc2736e2aeb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0656028 ([#11739](https://github.com/cloudquery/cloudquery/issues/11739)) ([7a6ad49](https://github.com/cloudquery/cloudquery/commit/7a6ad49e8402d51e914d6fdc444956c89db91ad3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 1e68c51 ([#11637](https://github.com/cloudquery/cloudquery/issues/11637)) ([46043bc](https://github.com/cloudquery/cloudquery/commit/46043bce410f86ba42390a6b190f9232fc2a1ded))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 43638cb ([#11672](https://github.com/cloudquery/cloudquery/issues/11672)) ([3c60bbb](https://github.com/cloudquery/cloudquery/commit/3c60bbbb0233b17f934583766938780745145864))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 4d76231 ([#11532](https://github.com/cloudquery/cloudquery/issues/11532)) ([6f04233](https://github.com/cloudquery/cloudquery/commit/6f042333acbd2506f7800ccb89a8c5cbfb7ad8d4))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8366a22 ([#11717](https://github.com/cloudquery/cloudquery/issues/11717)) ([8eeff5b](https://github.com/cloudquery/cloudquery/commit/8eeff5b17486d72845f830b99983f950fee7f5a0))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 95d3199 ([#11708](https://github.com/cloudquery/cloudquery/issues/11708)) ([03f214f](https://github.com/cloudquery/cloudquery/commit/03f214f3dfd719b74ce9eb698ba255a8cf7528c7))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b0832be ([#11651](https://github.com/cloudquery/cloudquery/issues/11651)) ([71e8c29](https://github.com/cloudquery/cloudquery/commit/71e8c29624494a3e1cd104e46266a610ce57c83c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d864719 ([#11611](https://github.com/cloudquery/cloudquery/issues/11611)) ([557a290](https://github.com/cloudquery/cloudquery/commit/557a2903af272b8e2e4c9eebb36e39cd8a41a805))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to df3b664 ([#11882](https://github.com/cloudquery/cloudquery/issues/11882)) ([9635b22](https://github.com/cloudquery/cloudquery/commit/9635b22b10a2cd9ca0f91819cffb7f4ba75dc2d9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f060192 ([#11730](https://github.com/cloudquery/cloudquery/issues/11730)) ([c7019c2](https://github.com/cloudquery/cloudquery/commit/c7019c26c311f29b66c90fc5d461a0daf71d191c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f0dffc6 ([#11689](https://github.com/cloudquery/cloudquery/issues/11689)) ([18ac0e9](https://github.com/cloudquery/cloudquery/commit/18ac0e9dbef31d06701f1f13d263ad840ac60c5e))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.1.0 ([#11665](https://github.com/cloudquery/cloudquery/issues/11665)) ([d8947c9](https://github.com/cloudquery/cloudquery/commit/d8947c9efa6ab8bf3952ad9d929e8ed81f2dea55))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.0 ([#11720](https://github.com/cloudquery/cloudquery/issues/11720)) ([7ef521d](https://github.com/cloudquery/cloudquery/commit/7ef521db1423c6f0de197b08c73adf22c896f999))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.1 ([#11722](https://github.com/cloudquery/cloudquery/issues/11722)) ([309be72](https://github.com/cloudquery/cloudquery/commit/309be7276d7de157013c281b6fb3934513898b3f))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.3 ([#11726](https://github.com/cloudquery/cloudquery/issues/11726)) ([f0ca611](https://github.com/cloudquery/cloudquery/commit/f0ca61195014bde707761a15efa27a92955b59db))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.4 ([#11753](https://github.com/cloudquery/cloudquery/issues/11753)) ([cd4fe1c](https://github.com/cloudquery/cloudquery/commit/cd4fe1c54f85f8511252bebd5671361618ddb0d3))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.5.0 ([#11850](https://github.com/cloudquery/cloudquery/issues/11850)) ([3255857](https://github.com/cloudquery/cloudquery/commit/3255857938bf16862d52491f5c2a8a0fa53faef0))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.6.0 ([#11916](https://github.com/cloudquery/cloudquery/issues/11916)) ([421e752](https://github.com/cloudquery/cloudquery/commit/421e7529360965175c8d156ff006d2b703ee9da2))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.6 ([#11473](https://github.com/cloudquery/cloudquery/issues/11473)) ([7272133](https://github.com/cloudquery/cloudquery/commit/72721336632e127dd37de4541f2f503bf4f73fb6))
* Github source pagination ([#11727](https://github.com/cloudquery/cloudquery/issues/11727)) ([f830ede](https://github.com/cloudquery/cloudquery/commit/f830ede1cf73c02985591bfcb41cbd2d81192d7f))

## [6.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v6.0.1...plugins-source-github-v6.0.2) (2023-06-13)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0f7bd3b ([#11412](https://github.com/cloudquery/cloudquery/issues/11412)) ([dd1e2e8](https://github.com/cloudquery/cloudquery/commit/dd1e2e892d95515fd7332339262abaefd2a256c5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 7f6aaff ([#11432](https://github.com/cloudquery/cloudquery/issues/11432)) ([55dfebc](https://github.com/cloudquery/cloudquery/commit/55dfebc064608fb47caaf3b8e68c8002de8a7dc3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8f72077 ([#11395](https://github.com/cloudquery/cloudquery/issues/11395)) ([d91fc5c](https://github.com/cloudquery/cloudquery/commit/d91fc5ce24f64c29fff6988b19ec2c2775cc379b))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 90670b8 ([#11279](https://github.com/cloudquery/cloudquery/issues/11279)) ([a6cdc91](https://github.com/cloudquery/cloudquery/commit/a6cdc912e4b38a3faf798c5147a986ffe2539643))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b359e74 ([#11405](https://github.com/cloudquery/cloudquery/issues/11405)) ([5d92765](https://github.com/cloudquery/cloudquery/commit/5d927659bd4f7c445a0e312487f1655ffb9a60f6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d8eacf8 ([#11449](https://github.com/cloudquery/cloudquery/issues/11449)) ([742dafd](https://github.com/cloudquery/cloudquery/commit/742dafd5bf5cdc8facb94fda5de1d84c88897cbd))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e258cfb ([#11391](https://github.com/cloudquery/cloudquery/issues/11391)) ([eacbe9a](https://github.com/cloudquery/cloudquery/commit/eacbe9ad3ea16d88f27c4593fa2774574ac8fe4e))

## [6.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v6.0.0...plugins-source-github-v6.0.1) (2023-06-06)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/beatlabs/github-auth digest to 0a8b9e6 ([#11152](https://github.com/cloudquery/cloudquery/issues/11152)) ([5505b5c](https://github.com/cloudquery/cloudquery/commit/5505b5cac50e73b310ed38b9e8488553c97e2389))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 20b0de9 ([#11199](https://github.com/cloudquery/cloudquery/issues/11199)) ([dc3565d](https://github.com/cloudquery/cloudquery/commit/dc3565d3fd6a640d9d10b4fd3a7fe6009a9d02a5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a7aad4c ([#11184](https://github.com/cloudquery/cloudquery/issues/11184)) ([8a0822e](https://github.com/cloudquery/cloudquery/commit/8a0822e31fc0eef99de2cdd2bd6d7e4c8b4131bf))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update golang.org/x/exp digest to 2e198f4 ([#11155](https://github.com/cloudquery/cloudquery/issues/11155)) ([c46c62b](https://github.com/cloudquery/cloudquery/commit/c46c62b68692f527485d7f4b84265abc5dc1142c))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.3 ([#11150](https://github.com/cloudquery/cloudquery/issues/11150)) ([dc00994](https://github.com/cloudquery/cloudquery/commit/dc00994e32936af7e9893c93561d0f9df225a929))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.4 ([#11244](https://github.com/cloudquery/cloudquery/issues/11244)) ([8fceef6](https://github.com/cloudquery/cloudquery/commit/8fceef6f9041e173923555d8ff221cfe83b424c2))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.7.0 ([#11113](https://github.com/cloudquery/cloudquery/issues/11113)) ([487bf87](https://github.com/cloudquery/cloudquery/commit/487bf871afe360cb8d9d592dfea48837d6e7cf27))

## [6.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v5.2.0...plugins-source-github-v6.0.0) (2023-05-29)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose). You will also need to update destinations depending on which one you use:
    - Azure Blob Storage >= v3.2.0
    - BigQuery >= v3.0.0
    - ClickHouse >= v3.1.1
    - DuckDB >= v1.1.6
    - Elasticsearch >= v2.0.0
    - File >= v3.2.0
    - Firehose >= v2.0.2
    - GCS >= v3.2.0
    - Gremlin >= v2.1.10
    - Kafka >= v3.0.1
    - Meilisearch >= v2.0.1
    - Microsoft SQL Server >= v4.2.0
    - MongoDB >= v2.0.1
    - MySQL >= v2.0.2
    - Neo4j >= v3.0.0
    - PostgreSQL >= v4.2.0
    - S3 >= v4.4.0
    - Snowflake >= v2.1.1
    - SQLite >= v2.2.0

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([#11008](https://github.com/cloudquery/cloudquery/issues/11008)) ([b00e212](https://github.com/cloudquery/cloudquery/commit/b00e212e024d1cf448de675ba657adcff902d900))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.6.7 ([#11043](https://github.com/cloudquery/cloudquery/issues/11043)) ([3c6d885](https://github.com/cloudquery/cloudquery/commit/3c6d885c3d201b0b39cbc1406c6e54a57ec5ed5f))
* **deps:** Update module github.com/gofri/go-github-ratelimit to v1.0.3 ([#11085](https://github.com/cloudquery/cloudquery/issues/11085)) ([4cadae8](https://github.com/cloudquery/cloudquery/commit/4cadae84e7f8949e768c38d0aa62b01d9d3c9ebf))

## [5.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v5.1.1...plugins-source-github-v5.2.0) (2023-05-16)


### Features

* Add GitHub Enterprise support on GitHub source ([#10776](https://github.com/cloudquery/cloudquery/issues/10776)) ([3bc226b](https://github.com/cloudquery/cloudquery/commit/3bc226bb7d704ee76f71566e9a043b95b164ce05))
* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))

## [5.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v5.1.0...plugins-source-github-v5.1.1) (2023-05-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update github.com/beatlabs/github-auth digest to f39a1f0 ([#10464](https://github.com/cloudquery/cloudquery/issues/10464)) ([683aba0](https://github.com/cloudquery/cloudquery/commit/683aba0cb9ee7f49502a488acccb44b6336503c2))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))

## [5.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v5.0.1...plugins-source-github-v5.1.0) (2023-04-25)


### This Release has the Following Changes to Tables
- Table `github_repository_branches` was added

### Features

* **github:** Add `github_repository_branches` table ([#10044](https://github.com/cloudquery/cloudquery/issues/10044)) ([c5eade6](https://github.com/cloudquery/cloudquery/commit/c5eade6105e9eeb30462b77f00121bf85db013d1))
* **github:** Upgrade to `github.com/cloudquery/plugin-sdk/v2` ([#10052](https://github.com/cloudquery/cloudquery/issues/10052)) ([44744e8](https://github.com/cloudquery/cloudquery/commit/44744e8d2c9674d25125ab99dd82f50da2afc58a)), closes [#10026](https://github.com/cloudquery/cloudquery/issues/10026)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.2.0 ([#10135](https://github.com/cloudquery/cloudquery/issues/10135)) ([cf33b89](https://github.com/cloudquery/cloudquery/commit/cf33b892ead0bb231e3956aa70967de552a21624))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.2.2 ([#10143](https://github.com/cloudquery/cloudquery/issues/10143)) ([8f887e0](https://github.com/cloudquery/cloudquery/commit/8f887e05de2096e8efd1e55863a8cf3c7620ccc3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.0 ([#10163](https://github.com/cloudquery/cloudquery/issues/10163)) ([9a7f214](https://github.com/cloudquery/cloudquery/commit/9a7f21460772200e7a588409ebc7eb19f97b195b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.1 ([#10175](https://github.com/cloudquery/cloudquery/issues/10175)) ([5b53423](https://github.com/cloudquery/cloudquery/commit/5b53423e72672f6c2bfb8ae00cfce1641410443e))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.3 ([#10187](https://github.com/cloudquery/cloudquery/issues/10187)) ([b185248](https://github.com/cloudquery/cloudquery/commit/b1852480b6ec8b721d94c72d8435051352f26932))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.4 ([#10196](https://github.com/cloudquery/cloudquery/issues/10196)) ([c6d2f59](https://github.com/cloudquery/cloudquery/commit/c6d2f59c7d77177a351cb82ecdc381dec6aad30c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.5 ([#10200](https://github.com/cloudquery/cloudquery/issues/10200)) ([5a33693](https://github.com/cloudquery/cloudquery/commit/5a33693fe29f7068b03d80be1859d6e479c42c0d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.6 ([#10208](https://github.com/cloudquery/cloudquery/issues/10208)) ([91c80a7](https://github.com/cloudquery/cloudquery/commit/91c80a795b46480447cfaef67c4db721a31e3206))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.8 ([#10213](https://github.com/cloudquery/cloudquery/issues/10213)) ([f358666](https://github.com/cloudquery/cloudquery/commit/f35866611cd206c37e6e9f9ad3329561e4cb32af))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.4.0 ([#10278](https://github.com/cloudquery/cloudquery/issues/10278)) ([a0a713e](https://github.com/cloudquery/cloudquery/commit/a0a713e8490b970b9d8bfaa1b50e01f43ff51c36))

## [5.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v5.0.0...plugins-source-github-v5.0.1) (2023-04-13)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))

## [5.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v4.1.3...plugins-source-github-v5.0.0) (2023-04-04)


### This Release has the Following Changes to Tables
- Table `github_issues`: column added with name `repository_id (PK)` and type `Int` (:warning: breaking)

### ⚠ BREAKING CHANGES

* **github:** Add `repository_id` to `github_issues` primary key ([#9615](https://github.com/cloudquery/cloudquery/issues/9615))

### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))
* **github:** Add `repository_id` to `github_issues` primary key ([#9615](https://github.com/cloudquery/cloudquery/issues/9615)) ([68625cf](https://github.com/cloudquery/cloudquery/commit/68625cf3e26c08f4ede65f696f308b4eee6d6e7d))
* **github:** Use cursor-based sync for `github_repository_dependabot_alerts` ([#9554](https://github.com/cloudquery/cloudquery/issues/9554)) ([feb4d13](https://github.com/cloudquery/cloudquery/commit/feb4d1318115aa2b31201041ab8f3337c33ece04))

## [4.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v4.1.2...plugins-source-github-v4.1.3) (2023-03-28)


### Bug Fixes

* **github:** Add pagination for dependabot resources ([#9461](https://github.com/cloudquery/cloudquery/issues/9461)) ([730f448](https://github.com/cloudquery/cloudquery/commit/730f4488da286687d844317509677aa7fb5a4822))

## [4.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v4.1.1...plugins-source-github-v4.1.2) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [4.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v4.1.0...plugins-source-github-v4.1.1) (2023-03-15)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))

## [4.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v4.0.1...plugins-source-github-v4.1.0) (2023-03-07)


### Features

* **docs:** Render tables as a part of the Website and add a [tables search box](https://www.cloudquery.io/tables). The equivalent of the GitHub README.md file is now under each plugin's docs section, for example https://www.cloudquery.io/docs/plugins/sources/aws/tables. The Website HTML page is built from the GitHub markdown file located under each plugin's path in our Website code, for example https://github.com/cloudquery/cloudquery/blob/main/website/pages/docs/plugins/sources/aws/tables.md. For the list of all plugins table files as they are stored on GitHub see https://github.com/cloudquery/cloudquery/tree/main/website/tables ([342b0c5](https://github.com/cloudquery/cloudquery/commit/342b0c569fd28ee26ea3e09ec6d787f85c49f16c))
* **github:** Add support for app-based authentication and filtering by repository ([#8556](https://github.com/cloudquery/cloudquery/issues/8556)) ([d5d5ddd](https://github.com/cloudquery/cloudquery/commit/d5d5ddd32c5835564c012df5e8856603d5a5cef6))


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to c95f2b4 ([#8560](https://github.com/cloudquery/cloudquery/issues/8560)) ([9c3bd5b](https://github.com/cloudquery/cloudquery/commit/9c3bd5b68f9741a360fde6c54bf3f5f3efe06d9e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))

## [4.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v4.0.0...plugins-source-github-v4.0.1) (2023-02-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))

## [4.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v3.0.1...plugins-source-github-v4.0.0) (2023-02-21)


### ⚠ BREAKING CHANGES

* **github:** Remove `repository` column from `github_workflows`
* **github:** Change `github_organization_dependabot_alerts` PK from `(org,number)` to `(org,html_url)` ([#8070](https://github.com/cloudquery/cloudquery/issues/8070))

### Features

* **github-resources:** Add Traffic resources ([#8085](https://github.com/cloudquery/cloudquery/issues/8085)) ([cc22f56](https://github.com/cloudquery/cloudquery/commit/cc22f566515bc23fae0104ca11c3dbefa1a2d17f))
* **github:** Add repositories multiplexer ([#8074](https://github.com/cloudquery/cloudquery/issues/8074)) ([f5f874d](https://github.com/cloudquery/cloudquery/commit/f5f874d5917ee4b9aeb3d2f0efef21bf5f492693))
* **github:** Handle secondary rate limit ([#8078](https://github.com/cloudquery/cloudquery/issues/8078)) ([bf1f1cc](https://github.com/cloudquery/cloudquery/commit/bf1f1cc890af13962617bb858594c123b90c24d3))
* **github:** Remove `repository` column from `github_workflows` ([21905a4](https://github.com/cloudquery/cloudquery/commit/21905a4a727023b20cab59387309cb70933fe5d7))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.38.2 ([#8156](https://github.com/cloudquery/cloudquery/issues/8156)) ([ac2d2d7](https://github.com/cloudquery/cloudquery/commit/ac2d2d70d5c4bc45fb8734bd4deb8a1e36074f6d))
* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))
* **github:** Add `repository_id` to `github_workflows` PK ([21905a4](https://github.com/cloudquery/cloudquery/commit/21905a4a727023b20cab59387309cb70933fe5d7))
* **github:** Change `github_organization_dependabot_alerts` PK from `(org,number)` to `(org,html_url)` ([#8070](https://github.com/cloudquery/cloudquery/issues/8070)) ([a7c64cd](https://github.com/cloudquery/cloudquery/commit/a7c64cd4dc946c1def29a32f61062decc958e1ce))

## [3.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v3.0.0...plugins-source-github-v3.0.1) (2023-02-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.1 ([#7930](https://github.com/cloudquery/cloudquery/issues/7930)) ([39dccc1](https://github.com/cloudquery/cloudquery/commit/39dccc1bf81f4eb02d181ba0c47b37038a4c5455))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.0 ([#7933](https://github.com/cloudquery/cloudquery/issues/7933)) ([dc9cffb](https://github.com/cloudquery/cloudquery/commit/dc9cffbf37bbc6fae73a20bf47e6bbf17e74d1f9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.1 ([#8008](https://github.com/cloudquery/cloudquery/issues/8008)) ([c47aac0](https://github.com/cloudquery/cloudquery/commit/c47aac0b5e3190a04299713651b97e360043911f))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v2.1.4...plugins-source-github-v3.0.0) (2023-02-07)


### ⚠ BREAKING CHANGES

* **github-resources:** The `request` and `response` columns for `github_hook_deliveries` are now of type `JSON` instead of `String`.

### Bug Fixes

* **deps:** Update golang.org/x/exp digest to f062dba ([#7531](https://github.com/cloudquery/cloudquery/issues/7531)) ([59d5575](https://github.com/cloudquery/cloudquery/commit/59d55758b0951553b8d246d1e78b4e3917ff1976))
* **deps:** Update google.golang.org/genproto digest to 1c01626 ([#7533](https://github.com/cloudquery/cloudquery/issues/7533)) ([c549c27](https://github.com/cloudquery/cloudquery/commit/c549c275077f1cdfb9df0b3f3c129cbf0b150552))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.0 ([#7595](https://github.com/cloudquery/cloudquery/issues/7595)) ([c5adc75](https://github.com/cloudquery/cloudquery/commit/c5adc750d4b0242563997c04c582f8da27913095))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))
* **github-resources:** Resolver Org Hooks request/response using GetHookDelivery ([52c6464](https://github.com/cloudquery/cloudquery/commit/52c6464865ae6a3d65ed53533c0b64d4d608c8a8))

## [2.1.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v2.1.3...plugins-source-github-v2.1.4) (2023-01-31)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))

## [2.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v2.1.2...plugins-source-github-v2.1.3) (2023-01-24)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))

## [2.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v2.1.1...plugins-source-github-v2.1.2) (2023-01-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.1 ([#6805](https://github.com/cloudquery/cloudquery/issues/6805)) ([9da0ce2](https://github.com/cloudquery/cloudquery/commit/9da0ce283f50410eb9274375ec1d22131a80d937))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.26.0 ([#6839](https://github.com/cloudquery/cloudquery/issues/6839)) ([6ccda8d](https://github.com/cloudquery/cloudquery/commit/6ccda8d0bc6e7ce75f4a64a18911e349ccaac277))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.27.0 ([#6856](https://github.com/cloudquery/cloudquery/issues/6856)) ([545799b](https://github.com/cloudquery/cloudquery/commit/545799bb0481087e187b5f27c88f5dde9c99f2f0))

## [2.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v2.1.0...plugins-source-github-v2.1.1) (2023-01-12)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.2 ([#6695](https://github.com/cloudquery/cloudquery/issues/6695)) ([694ab9f](https://github.com/cloudquery/cloudquery/commit/694ab9f3e20473146e3620d7b03bb17eb259d697))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.0 ([#6745](https://github.com/cloudquery/cloudquery/issues/6745)) ([9c41854](https://github.com/cloudquery/cloudquery/commit/9c418547c3bbff97449765e337182230fb5e40d5))
* **github:** Fix panic in workflow fetching ([#6693](https://github.com/cloudquery/cloudquery/issues/6693)) ([a582f0b](https://github.com/cloudquery/cloudquery/commit/a582f0b76bb28be4ca6a9448fcf452d7408b12fd))

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v2.0.0...plugins-source-github-v2.1.0) (2023-01-10)


### Features

* **github:** Migrate codegen to transformations ([#6428](https://github.com/cloudquery/cloudquery/issues/6428)) ([71c84fc](https://github.com/cloudquery/cloudquery/commit/71c84fcee0ae1496f16d290b605bda81316f3087))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.2 ([#6260](https://github.com/cloudquery/cloudquery/issues/6260)) ([805972a](https://github.com/cloudquery/cloudquery/commit/805972aa67ce54e3358501c6b7ee5d85e5f65cac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.18.0 ([#6339](https://github.com/cloudquery/cloudquery/issues/6339)) ([158365a](https://github.com/cloudquery/cloudquery/commit/158365a78dfa4389074f716a0f581f18fedc1080))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.19.0 ([#6363](https://github.com/cloudquery/cloudquery/issues/6363)) ([ae6967c](https://github.com/cloudquery/cloudquery/commit/ae6967c22002c554a083f444eb611ac3e6d2698f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.20.0 ([#6376](https://github.com/cloudquery/cloudquery/issues/6376)) ([d6187ec](https://github.com/cloudquery/cloudquery/commit/d6187ec584f13be4fe9362dd393385b19d386113))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.21.0 ([#6382](https://github.com/cloudquery/cloudquery/issues/6382)) ([5baea40](https://github.com/cloudquery/cloudquery/commit/5baea40d2aec4e807db839c928be2e037d572bef))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.22.0 ([#6516](https://github.com/cloudquery/cloudquery/issues/6516)) ([b7e4e73](https://github.com/cloudquery/cloudquery/commit/b7e4e737a5f4d8f254960426ea8ba555d8f9b944))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.23.0 ([#6522](https://github.com/cloudquery/cloudquery/issues/6522)) ([ce24f1d](https://github.com/cloudquery/cloudquery/commit/ce24f1d64394cbb5ab07dcaa4af66c53f77f700f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.1 ([#6553](https://github.com/cloudquery/cloudquery/issues/6553)) ([392b848](https://github.com/cloudquery/cloudquery/commit/392b848b3124f9cf28f6234fdb9a43d671069879))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.4.2...plugins-source-github-v2.0.0) (2023-01-03)


### ⚠ BREAKING CHANGES

* **github:** Reorder PK columns in for `github_hook_deliveries`, `github_team_members` and `github_team_repositories`

### Features

* **github:** Add Releases and Release Assets ([#6097](https://github.com/cloudquery/cloudquery/issues/6097)) ([750ff36](https://github.com/cloudquery/cloudquery/commit/750ff360cc18600557eb8ff7b66f0a39cd5c233e))
* **github:** Add tables related to GitHub Dependabot API ([#6046](https://github.com/cloudquery/cloudquery/issues/6046)) ([dd0592b](https://github.com/cloudquery/cloudquery/commit/dd0592bc985f4613c190904d4e4d21a94f940f4e))
* **github:** Update GitHub client ([5659f86](https://github.com/cloudquery/cloudquery/commit/5659f8606859996728d5857a57f87013abd4ee03))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.14.0 ([#6025](https://github.com/cloudquery/cloudquery/issues/6025)) ([35b2cfc](https://github.com/cloudquery/cloudquery/commit/35b2cfc7fc7bcdaceb7ee674e3a17f0f5673b366))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.0 ([#6071](https://github.com/cloudquery/cloudquery/issues/6071)) ([684b525](https://github.com/cloudquery/cloudquery/commit/684b525aaa285fcae70dd87af56679c1205adebe))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.1 ([#6079](https://github.com/cloudquery/cloudquery/issues/6079)) ([650659c](https://github.com/cloudquery/cloudquery/commit/650659c3c6766df571868e2ec3a2007cb76696eb))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.0 ([#6098](https://github.com/cloudquery/cloudquery/issues/6098)) ([7bacdf3](https://github.com/cloudquery/cloudquery/commit/7bacdf3364716eab08fa1a84ae4047b42edeee7e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.1 ([#6214](https://github.com/cloudquery/cloudquery/issues/6214)) ([53b2415](https://github.com/cloudquery/cloudquery/commit/53b241508d7511d4b5fa74cc4262d180c1e6df66))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.0 ([#6256](https://github.com/cloudquery/cloudquery/issues/6256)) ([b19f6cd](https://github.com/cloudquery/cloudquery/commit/b19f6cd8e2c39994aeb19d78e78e927d6c3cf580))
* **github:** Change PK columns order for some resources ([5659f86](https://github.com/cloudquery/cloudquery/commit/5659f8606859996728d5857a57f87013abd4ee03))

## [1.4.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.4.1...plugins-source-github-v1.4.2) (2022-12-27)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.13.1 ([#5897](https://github.com/cloudquery/cloudquery/issues/5897)) ([ad15915](https://github.com/cloudquery/cloudquery/commit/ad15915f2951a75729859f6f1377ed789f8ba115))

## [1.4.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.4.0...plugins-source-github-v1.4.1) (2022-12-20)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.2 ([#5583](https://github.com/cloudquery/cloudquery/issues/5583)) ([d721c4e](https://github.com/cloudquery/cloudquery/commit/d721c4e06b8a97b5373215aca0e4ed64942ac489))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.3 ([#5639](https://github.com/cloudquery/cloudquery/issues/5639)) ([6452d0e](https://github.com/cloudquery/cloudquery/commit/6452d0ed5a44abad9d7530af6e79cde6504d0c4c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.4 ([#5649](https://github.com/cloudquery/cloudquery/issues/5649)) ([b4aa889](https://github.com/cloudquery/cloudquery/commit/b4aa889e396db3b0887d1684e4bc07da6050af43))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.5 ([#5661](https://github.com/cloudquery/cloudquery/issues/5661)) ([b354b8a](https://github.com/cloudquery/cloudquery/commit/b354b8a3683fa2bc918c1002afac487427d65a5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.6 ([#5790](https://github.com/cloudquery/cloudquery/issues/5790)) ([8e2663c](https://github.com/cloudquery/cloudquery/commit/8e2663c17c3347afd5e53f665462adc3e709c96c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.7 ([#5797](https://github.com/cloudquery/cloudquery/issues/5797)) ([15da529](https://github.com/cloudquery/cloudquery/commit/15da5294786fa2656228ca5bbc48ef1fc44e486b))

## [1.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.3.0...plugins-source-github-v1.4.0) (2022-12-13)


### Features

* Add GH Workflow support ([#3195](https://github.com/cloudquery/cloudquery/issues/3195)) ([098ac50](https://github.com/cloudquery/cloudquery/commit/098ac501ab7104876e1625d1d30e2213a3542551))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.1 ([#5458](https://github.com/cloudquery/cloudquery/issues/5458)) ([58b7432](https://github.com/cloudquery/cloudquery/commit/58b74321cd253c9a843c8c103f324abb93952195))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.2 ([#5497](https://github.com/cloudquery/cloudquery/issues/5497)) ([c1876cf](https://github.com/cloudquery/cloudquery/commit/c1876cf793b43d825a25fb3c9ba4996e4b09964f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.0 ([#5539](https://github.com/cloudquery/cloudquery/issues/5539)) ([fb71293](https://github.com/cloudquery/cloudquery/commit/fb71293d5cfe1b2ef32ba83d604ac3c48e662bce))

## [1.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.19...plugins-source-github-v1.3.0) (2022-12-06)


### Features

* **website:** Add plugins tables ([#5259](https://github.com/cloudquery/cloudquery/issues/5259)) ([c336f4e](https://github.com/cloudquery/cloudquery/commit/c336f4e25e192ffdd4c211d4a35b67b71d01d1f8))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.0 ([#5416](https://github.com/cloudquery/cloudquery/issues/5416)) ([2e7ca35](https://github.com/cloudquery/cloudquery/commit/2e7ca35922fdb14fd717f582aaaa9693dae2ef4c))
* **deps:** Update module github.com/google/go-github/v45 to v48 ([#5243](https://github.com/cloudquery/cloudquery/issues/5243)) ([0a2c74a](https://github.com/cloudquery/cloudquery/commit/0a2c74a9d2f586a0ff8d4b0ce786d2ee22445f74))
* **deps:** Update module golang.org/x/oauth2 to v0.2.0 ([#5229](https://github.com/cloudquery/cloudquery/issues/5229)) ([7fed2bd](https://github.com/cloudquery/cloudquery/commit/7fed2bdbf229109e1bc6560eb8e94d2ad5ae03d9))

## [1.2.19](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.18...plugins-source-github-v1.2.19) (2022-11-30)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.10.0 ([#5153](https://github.com/cloudquery/cloudquery/issues/5153)) ([ea1f77e](https://github.com/cloudquery/cloudquery/commit/ea1f77e910f430287600e74cedd7d3f4ae79eb18))
* **deps:** Update plugin-sdk for github to v1.8.1 ([#5039](https://github.com/cloudquery/cloudquery/issues/5039)) ([3ddda50](https://github.com/cloudquery/cloudquery/commit/3ddda50a722f959a91283bb359ea448e9fb10ac2))
* **deps:** Update plugin-sdk for github to v1.9.0 ([#5081](https://github.com/cloudquery/cloudquery/issues/5081)) ([33e2dcc](https://github.com/cloudquery/cloudquery/commit/33e2dcc6c826773d3bb47aed98596be793be7427))

## [1.2.18](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.17...plugins-source-github-v1.2.18) (2022-11-23)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.7.0 ([#4910](https://github.com/cloudquery/cloudquery/issues/4910)) ([faf723b](https://github.com/cloudquery/cloudquery/commit/faf723b18bb808f7859443df7f1ae24478210cf1))
* **deps:** Update plugin-sdk for github to v1.8.0 ([#4973](https://github.com/cloudquery/cloudquery/issues/4973)) ([f9c16d0](https://github.com/cloudquery/cloudquery/commit/f9c16d0d6749bb8c3aef4e3160983b524657e4a6))

## [1.2.17](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.16...plugins-source-github-v1.2.17) (2022-11-21)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.6.0 ([#4848](https://github.com/cloudquery/cloudquery/issues/4848)) ([8d59baf](https://github.com/cloudquery/cloudquery/commit/8d59baff0a3a6e40617400e72a82a1068d2032f0))

## [1.2.16](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.15...plugins-source-github-v1.2.16) (2022-11-15)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.5.3 ([#4647](https://github.com/cloudquery/cloudquery/issues/4647)) ([262b0e0](https://github.com/cloudquery/cloudquery/commit/262b0e0788ac76c120eb69609467900ac021dbfc))

## [1.2.15](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.14...plugins-source-github-v1.2.15) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.5.2 ([#4562](https://github.com/cloudquery/cloudquery/issues/4562)) ([7258b4e](https://github.com/cloudquery/cloudquery/commit/7258b4eb8906355e99ba07d1de4acac587e599b3))

## [1.2.14](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.13...plugins-source-github-v1.2.14) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.5.1 ([#4503](https://github.com/cloudquery/cloudquery/issues/4503)) ([a800549](https://github.com/cloudquery/cloudquery/commit/a8005494ae203a71d2a571ee33d67749e5a5b320))

## [1.2.13](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.12...plugins-source-github-v1.2.13) (2022-11-11)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.5.0 ([#4392](https://github.com/cloudquery/cloudquery/issues/4392)) ([639ee30](https://github.com/cloudquery/cloudquery/commit/639ee306228f72bdf453aa545c8be5611a592db7))

## [1.2.12](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.11...plugins-source-github-v1.2.12) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.4.1 ([#4295](https://github.com/cloudquery/cloudquery/issues/4295)) ([8b3faf7](https://github.com/cloudquery/cloudquery/commit/8b3faf7f29206eccb714628e22fd2130813bce06))

## [1.2.11](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.10...plugins-source-github-v1.2.11) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.4.0 ([#4233](https://github.com/cloudquery/cloudquery/issues/4233)) ([87b23a3](https://github.com/cloudquery/cloudquery/commit/87b23a3a4ee8b396e31d3fd965bc2766944facce))

## [1.2.10](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.9...plugins-source-github-v1.2.10) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.3.2 ([#4200](https://github.com/cloudquery/cloudquery/issues/4200)) ([9c7b2b6](https://github.com/cloudquery/cloudquery/commit/9c7b2b6c42bf227527f2e4e0109cca768aa4b85b))

## [1.2.9](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.8...plugins-source-github-v1.2.9) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.3.1 ([#4150](https://github.com/cloudquery/cloudquery/issues/4150)) ([dc0c2af](https://github.com/cloudquery/cloudquery/commit/dc0c2afeeb0c9a250785db92dbc10604baa7d648))

## [1.2.8](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.7...plugins-source-github-v1.2.8) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.3.0 ([#4075](https://github.com/cloudquery/cloudquery/issues/4075)) ([91776f4](https://github.com/cloudquery/cloudquery/commit/91776f43eea99f61c961fdff53ee9d017f4773ad))

## [1.2.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.6...plugins-source-github-v1.2.7) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.2.0 ([#4043](https://github.com/cloudquery/cloudquery/issues/4043)) ([56ebf5b](https://github.com/cloudquery/cloudquery/commit/56ebf5bbd61c2255841b7ddf1e7cf889bb139890))

## [1.2.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.5...plugins-source-github-v1.2.6) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for csv to v1.1.0 ([#3918](https://github.com/cloudquery/cloudquery/issues/3918)) ([f1acd68](https://github.com/cloudquery/cloudquery/commit/f1acd688fcd90011cc9be1be2285e3fe9369e341))

## [1.2.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.4...plugins-source-github-v1.2.5) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1.0.3 ([#3852](https://github.com/cloudquery/cloudquery/issues/3852)) ([0f1d51d](https://github.com/cloudquery/cloudquery/commit/0f1d51d8630de73780877c13ee226337c80c80ad))
* **deps:** Update plugin-sdk for github to v1.1.0 ([#3921](https://github.com/cloudquery/cloudquery/issues/3921)) ([14b4922](https://github.com/cloudquery/cloudquery/commit/14b49227ec87013d538a5f195329e5cb1d34d64d))
* **deps:** Upgrade plugin-sdk to v1.0.4 for plugins ([#3889](https://github.com/cloudquery/cloudquery/issues/3889)) ([6767243](https://github.com/cloudquery/cloudquery/commit/6767243ec70bfae7a4c457bf4b5edf013c54c392))

## [1.2.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.3...plugins-source-github-v1.2.4) (2022-11-07)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v1 ([#3792](https://github.com/cloudquery/cloudquery/issues/3792)) ([d180e25](https://github.com/cloudquery/cloudquery/commit/d180e25d2a0f140292879cba2b882817015b9941))

## [1.2.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.2...plugins-source-github-v1.2.3) (2022-11-07)


### Bug Fixes

* **deps:** Update SDK to v0.13.23 ([#3738](https://github.com/cloudquery/cloudquery/issues/3738)) ([a56b65c](https://github.com/cloudquery/cloudquery/commit/a56b65c811e083ab38c89bd345aed4e4520ec12a))

## [1.2.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.1...plugins-source-github-v1.2.2) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.13.20 ([#3576](https://github.com/cloudquery/cloudquery/issues/3576)) ([58d16c2](https://github.com/cloudquery/cloudquery/commit/58d16c2579dccf8909bb5447910fefb1d03f1066))

## [1.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.2.0...plugins-source-github-v1.2.1) (2022-11-03)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.13.19 ([#3507](https://github.com/cloudquery/cloudquery/issues/3507)) ([8ccdef6](https://github.com/cloudquery/cloudquery/commit/8ccdef6725c0809a6682b546bb9263c9a694cae8))

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.1.0...plugins-source-github-v1.2.0) (2022-11-01)


### Features

* Migrate cli, plugins and destinations to new type system ([#3323](https://github.com/cloudquery/cloudquery/issues/3323)) ([f265a94](https://github.com/cloudquery/cloudquery/commit/f265a94448ad55c968b26ba8a19681bc81086c11))


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.13.18 ([#3405](https://github.com/cloudquery/cloudquery/issues/3405)) ([b6be94c](https://github.com/cloudquery/cloudquery/commit/b6be94c1c9d3a75d2cc6967eb5d7dd7a19abbc51))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.9...plugins-source-github-v1.1.0) (2022-10-31)


### Features

* Update all plugins to SDK with metrics and DFS scheduler ([#3286](https://github.com/cloudquery/cloudquery/issues/3286)) ([a35b8e8](https://github.com/cloudquery/cloudquery/commit/a35b8e89d625287a9b9406ff18cfac78ffdb1241))

## [1.0.9](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.8...plugins-source-github-v1.0.9) (2022-10-27)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.13.13 ([#3217](https://github.com/cloudquery/cloudquery/issues/3217)) ([269c0ac](https://github.com/cloudquery/cloudquery/commit/269c0ac546f569e8333f869ff0bbcc9dfcf75f37))
* **deps:** Update plugin-sdk for github to v0.13.14 ([#3233](https://github.com/cloudquery/cloudquery/issues/3233)) ([af8abb1](https://github.com/cloudquery/cloudquery/commit/af8abb10b33e90953958fb5a7498836b4f0d3211))

## [1.0.8](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.7...plugins-source-github-v1.0.8) (2022-10-20)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.13.12 ([#3105](https://github.com/cloudquery/cloudquery/issues/3105)) ([8146995](https://github.com/cloudquery/cloudquery/commit/81469952373813ee38b6973b7f29944b2efd15ff))

## [1.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.6...plugins-source-github-v1.0.7) (2022-10-19)


### Bug Fixes

* **deps:** Update plugin-sdk to v0.13.11 ([#3030](https://github.com/cloudquery/cloudquery/issues/3030)) ([9909c4a](https://github.com/cloudquery/cloudquery/commit/9909c4a0715a06b7c1d69c9bd23c500ac7b4adc1))

## [1.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.5...plugins-source-github-v1.0.6) (2022-10-18)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.13.9 ([#2932](https://github.com/cloudquery/cloudquery/issues/2932)) ([f798f4a](https://github.com/cloudquery/cloudquery/commit/f798f4a29a500f24701a98784da07c5f29296c60))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.4...plugins-source-github-v1.0.5) (2022-10-14)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.13.8 ([#2854](https://github.com/cloudquery/cloudquery/issues/2854)) ([cedb32b](https://github.com/cloudquery/cloudquery/commit/cedb32bbc832c4880d5cc7f15f1b974615563f90))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.3...plugins-source-github-v1.0.4) (2022-10-13)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.13.7 ([#2784](https://github.com/cloudquery/cloudquery/issues/2784)) ([ab5c43f](https://github.com/cloudquery/cloudquery/commit/ab5c43f955082402d09bda1f5529b434102bff0c))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.2...plugins-source-github-v1.0.3) (2022-10-12)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.13.6 ([#2723](https://github.com/cloudquery/cloudquery/issues/2723)) ([f98ee9c](https://github.com/cloudquery/cloudquery/commit/f98ee9c6dbe510e1717401ea0eeefbb514cc0567))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.1...plugins-source-github-v1.0.2) (2022-10-12)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.12.10 ([#2550](https://github.com/cloudquery/cloudquery/issues/2550)) ([7d92b69](https://github.com/cloudquery/cloudquery/commit/7d92b69f44bb8618ac72de1bf510f164e0b5cd91))
* Upgrade source SDK versions to v0.13.5 ([#2610](https://github.com/cloudquery/cloudquery/issues/2610)) ([611868e](https://github.com/cloudquery/cloudquery/commit/611868e7fbb707b524ccc5c04a7ff95fe122ae05))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v1.0.0...plugins-source-github-v1.0.1) (2022-10-09)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.12.3 ([#2358](https://github.com/cloudquery/cloudquery/issues/2358)) ([435db2c](https://github.com/cloudquery/cloudquery/commit/435db2c72be475147dd115f7d5444cde4e73ca00))
* **deps:** Update plugin-sdk for github to v0.12.4 ([#2400](https://github.com/cloudquery/cloudquery/issues/2400)) ([c38e624](https://github.com/cloudquery/cloudquery/commit/c38e6242b59598e4fc2ac6448de05c1894c26347))
* **deps:** Update plugin-sdk for github to v0.12.5 ([#2422](https://github.com/cloudquery/cloudquery/issues/2422)) ([7f4eb34](https://github.com/cloudquery/cloudquery/commit/7f4eb3416acfc23cb2b5443da7da03b70ae0fdba))
* **deps:** Update plugin-sdk for github to v0.12.6 ([#2438](https://github.com/cloudquery/cloudquery/issues/2438)) ([2b593d4](https://github.com/cloudquery/cloudquery/commit/2b593d4989a55069ff716c344a3dfd316f72d325))
* **deps:** Update plugin-sdk for github to v0.12.7 ([#2451](https://github.com/cloudquery/cloudquery/issues/2451)) ([1520961](https://github.com/cloudquery/cloudquery/commit/15209614ad1f14133d49bcf9fc13e554a1d5f340))
* **deps:** Update plugin-sdk for github to v0.12.8 ([#2501](https://github.com/cloudquery/cloudquery/issues/2501)) ([a9b8b36](https://github.com/cloudquery/cloudquery/commit/a9b8b36ee9b00309d58f646f751e4f27a58dfa46))
* **deps:** Update plugin-sdk for github to v0.12.9 ([#2515](https://github.com/cloudquery/cloudquery/issues/2515)) ([820484e](https://github.com/cloudquery/cloudquery/commit/820484e8ec362d4c79a39efb044ea57a5e0d4d25))

## [1.0.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/github/v0.1.10...plugins-source-github-v1.0.0) (2022-10-04)


### ⚠ BREAKING CHANGES

* [Official v1 release](https://www.cloudquery.io/blog/cloudquery-v1-release)

### Features

* [Official v1 release](https://www.cloudquery.io/blog/cloudquery-v1-release)

## [0.1.16-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v0.1.15-pre.0...plugins-source-github-v0.1.16-pre.0) (2022-10-03)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.11.6 ([#2257](https://github.com/cloudquery/cloudquery/issues/2257)) ([b0a3b33](https://github.com/cloudquery/cloudquery/commit/b0a3b331948b3e90fe810dc002519d690282dcd3))

## [0.1.15-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v0.1.14-pre.0...plugins-source-github-v0.1.15-pre.0) (2022-10-03)


### Bug Fixes

* **deps:** Update plugin-sdk for github to v0.11.5 ([#2231](https://github.com/cloudquery/cloudquery/issues/2231)) ([25b9349](https://github.com/cloudquery/cloudquery/commit/25b9349a2fc2908e22a93b8b5f4f57e92adafe9f))

## [0.1.14-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v0.1.13-pre.0...plugins-source-github-v0.1.14-pre.0) (2022-10-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.10.2 ([#2048](https://github.com/cloudquery/cloudquery/issues/2048)) ([e407991](https://github.com/cloudquery/cloudquery/commit/e4079914772d8191639b9935aa5970b8e27b082f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.0 ([#2135](https://github.com/cloudquery/cloudquery/issues/2135)) ([1729467](https://github.com/cloudquery/cloudquery/commit/1729467b2119555e18b15d73c91cd501ccf7ecb8))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.2 ([#2162](https://github.com/cloudquery/cloudquery/issues/2162)) ([5701aa5](https://github.com/cloudquery/cloudquery/commit/5701aa5b0a8d04e9e99e3efe6e27d5f7ff29b216))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.4 ([#2214](https://github.com/cloudquery/cloudquery/issues/2214)) ([aa82b42](https://github.com/cloudquery/cloudquery/commit/aa82b427f2b296589608b16472cb9b36f9f3818b))

## [0.1.13-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v0.1.12-pre.0...plugins-source-github-v0.1.13-pre.0) (2022-10-02)


### Bug Fixes

* Add Github table docs back ([#1987](https://github.com/cloudquery/cloudquery/issues/1987)) ([1a21b5a](https://github.com/cloudquery/cloudquery/commit/1a21b5a118ff8bf03f3d93a52a92d2d47c89c140))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.10.2 ([#2048](https://github.com/cloudquery/cloudquery/issues/2048)) ([e407991](https://github.com/cloudquery/cloudquery/commit/e4079914772d8191639b9935aa5970b8e27b082f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.0 ([#2135](https://github.com/cloudquery/cloudquery/issues/2135)) ([1729467](https://github.com/cloudquery/cloudquery/commit/1729467b2119555e18b15d73c91cd501ccf7ecb8))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.2 ([#2162](https://github.com/cloudquery/cloudquery/issues/2162)) ([5701aa5](https://github.com/cloudquery/cloudquery/commit/5701aa5b0a8d04e9e99e3efe6e27d5f7ff29b216))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.4 ([#2214](https://github.com/cloudquery/cloudquery/issues/2214)) ([aa82b42](https://github.com/cloudquery/cloudquery/commit/aa82b427f2b296589608b16472cb9b36f9f3818b))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.0 ([#1997](https://github.com/cloudquery/cloudquery/issues/1997)) ([4fa40da](https://github.com/cloudquery/cloudquery/commit/4fa40da04b427f864d2dc11f133e5c83e53ce4b6))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.1 ([#2024](https://github.com/cloudquery/cloudquery/issues/2024)) ([8f88de4](https://github.com/cloudquery/cloudquery/commit/8f88de4b4eaeabae7369ba309e765a252392ee8c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.2 ([#2044](https://github.com/cloudquery/cloudquery/issues/2044)) ([9b69b46](https://github.com/cloudquery/cloudquery/commit/9b69b468536521b20b77ec1fc180fc85aeeba376))

## [0.1.13-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-github-v0.1.12-pre.0...plugins-source-github-v0.1.13-pre.0) (2022-09-26)


### Bug Fixes

* Add Github table docs back ([#1987](https://github.com/cloudquery/cloudquery/issues/1987)) ([1a21b5a](https://github.com/cloudquery/cloudquery/commit/1a21b5a118ff8bf03f3d93a52a92d2d47c89c140))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.0 ([#1997](https://github.com/cloudquery/cloudquery/issues/1997)) ([4fa40da](https://github.com/cloudquery/cloudquery/commit/4fa40da04b427f864d2dc11f133e5c83e53ce4b6))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.1 ([#2024](https://github.com/cloudquery/cloudquery/issues/2024)) ([8f88de4](https://github.com/cloudquery/cloudquery/commit/8f88de4b4eaeabae7369ba309e765a252392ee8c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.2 ([#2044](https://github.com/cloudquery/cloudquery/issues/2044)) ([9b69b46](https://github.com/cloudquery/cloudquery/commit/9b69b468536521b20b77ec1fc180fc85aeeba376))

## [0.1.12-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/github/v0.1.11-pre.0...plugins/source/github/v0.1.12-pre.0) (2022-09-22)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.12 ([#1916](https://github.com/cloudquery/cloudquery/issues/1916)) ([27d8153](https://github.com/cloudquery/cloudquery/commit/27d81534baaa1312a6bd87294d298dd8b5348a79))

## [0.1.11-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/github-v0.1.10-pre.0...plugins/source/github/v0.1.11-pre.0) (2022-09-21)


### Features

* Migrate GitHub plugin to v2 ([#1776](https://github.com/cloudquery/cloudquery/issues/1776)) ([217609b](https://github.com/cloudquery/cloudquery/commit/217609b8504affb4c560a14a56b84e91f91fd426))


### Bug Fixes

* **deps:** Update golang.org/x/oauth2 digest to 0ebed06 ([#1651](https://github.com/cloudquery/cloudquery/issues/1651)) ([c31b001](https://github.com/cloudquery/cloudquery/commit/c31b00145b70316ca38873769e33b21db783b937))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.10 ([#1474](https://github.com/cloudquery/cloudquery/issues/1474)) ([b142e13](https://github.com/cloudquery/cloudquery/commit/b142e135172b1eed1abb2cbec85054ea7f66199d))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.11 ([#1491](https://github.com/cloudquery/cloudquery/issues/1491)) ([5140bef](https://github.com/cloudquery/cloudquery/commit/5140bef4aa7c50a97a604db1e92df75ead2893fc))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.12 ([#1503](https://github.com/cloudquery/cloudquery/issues/1503)) ([a740719](https://github.com/cloudquery/cloudquery/commit/a7407199c9617784a1834b9d0c42788e03301de5))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.9 ([#1286](https://github.com/cloudquery/cloudquery/issues/1286)) ([67ac422](https://github.com/cloudquery/cloudquery/commit/67ac422f392387e674cb70386e612befa5b455f0))
* Validate github config ([#1438](https://github.com/cloudquery/cloudquery/issues/1438)) ([8d37fc4](https://github.com/cloudquery/cloudquery/commit/8d37fc4a95a327bc9b83fa627a67ac5167e2d513))

## [0.1.10](https://github.com/cloudquery/cloudquery/compare/plugins/source/github/v0.1.9...plugins/source/github/v0.1.10) (2022-09-01)


### Bug Fixes

* **deps:** Update golang.org/x/oauth2 digest to 0ebed06 ([#1651](https://github.com/cloudquery/cloudquery/issues/1651)) ([c31b001](https://github.com/cloudquery/cloudquery/commit/c31b00145b70316ca38873769e33b21db783b937))

## [0.1.9](https://github.com/cloudquery/cloudquery/compare/plugins/source/github/v0.1.8...plugins/source/github/v0.1.9) (2022-08-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.12 ([#1503](https://github.com/cloudquery/cloudquery/issues/1503)) ([a740719](https://github.com/cloudquery/cloudquery/commit/a7407199c9617784a1834b9d0c42788e03301de5))

## [0.1.8](https://github.com/cloudquery/cloudquery/compare/plugins/source/github/v0.1.7...plugins/source/github/v0.1.8) (2022-08-18)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.11 ([#1491](https://github.com/cloudquery/cloudquery/issues/1491)) ([5140bef](https://github.com/cloudquery/cloudquery/commit/5140bef4aa7c50a97a604db1e92df75ead2893fc))

## [0.1.7](https://github.com/cloudquery/cloudquery/compare/plugins/source/github/v0.1.6...plugins/source/github/v0.1.7) (2022-08-18)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.10 ([#1474](https://github.com/cloudquery/cloudquery/issues/1474)) ([b142e13](https://github.com/cloudquery/cloudquery/commit/b142e135172b1eed1abb2cbec85054ea7f66199d))

## [0.1.6](https://github.com/cloudquery/cloudquery/compare/plugins/source/github/v0.1.5...plugins/source/github/v0.1.6) (2022-08-17)


### Bug Fixes

* Validate github config ([#1438](https://github.com/cloudquery/cloudquery/issues/1438)) ([8d37fc4](https://github.com/cloudquery/cloudquery/commit/8d37fc4a95a327bc9b83fa627a67ac5167e2d513))

## [0.1.5](https://github.com/cloudquery/cloudquery/compare/plugins/source/github-v0.1.4...plugins/source/github/v0.1.5) (2022-08-15)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.9 ([#1286](https://github.com/cloudquery/cloudquery/issues/1286)) ([67ac422](https://github.com/cloudquery/cloudquery/commit/67ac422f392387e674cb70386e612befa5b455f0))

## [0.1.4](https://github.com/cloudquery/cq-provider-github/compare/v0.1.3...v0.1.4) (2022-08-10)


### Features

* Ignore forbidden errors ([#27](https://github.com/cloudquery/cq-provider-github/issues/27)) ([535898a](https://github.com/cloudquery/cq-provider-github/commit/535898a020997b01db48d291cad56b275d1357bd))


### Bug Fixes

* Small readme fixes ([#25](https://github.com/cloudquery/cq-provider-github/issues/25)) ([4caea14](https://github.com/cloudquery/cq-provider-github/commit/4caea14916b0ccfd94083761f3233a9fa1ba6c8c))

## [0.1.3](https://github.com/cloudquery/cq-provider-github/compare/v0.1.2...v0.1.3) (2022-08-09)


### Bug Fixes

* Remove Imposter releaser ([#23](https://github.com/cloudquery/cq-provider-github/issues/23)) ([9d47435](https://github.com/cloudquery/cq-provider-github/commit/9d47435c982e12f4415ba2bc36059b4f8576d903))

## [0.1.2](https://github.com/cloudquery/cq-provider-github/compare/v0.1.1...v0.1.2) (2022-08-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.7 ([#15](https://github.com/cloudquery/cq-provider-github/issues/15)) ([2732a29](https://github.com/cloudquery/cq-provider-github/commit/2732a29fe4df57d0be0c9315cbc2a610c27c1eaf))

## [0.1.1](https://github.com/cloudquery/cq-provider-github/compare/v0.1.0...v0.1.1) (2022-08-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.7 ([#14](https://github.com/cloudquery/cq-provider-github/issues/14)) ([58d59db](https://github.com/cloudquery/cq-provider-github/commit/58d59db48a16c57c2ae3625b9140d4dc92f3a9b6))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.8 ([#19](https://github.com/cloudquery/cq-provider-github/issues/19)) ([36b0433](https://github.com/cloudquery/cq-provider-github/commit/36b04332a3b33a2422cb449aa9a4efcf08b6a68d))
* Issue fetching ([#20](https://github.com/cloudquery/cq-provider-github/issues/20)) ([1753dd6](https://github.com/cloudquery/cq-provider-github/commit/1753dd6caebd0a711fd9685b2f5ad7c101715402))

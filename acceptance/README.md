# Tests description for acceptance

## `appcharts_test.go`

- **Describe:** apps chart, LAppchart
    - **It:** lists the known app charts
    - **It:** shows the details of the standard app chart
    - **It:** shows the details of the custom chart
    - **It:** fails to show the details of a bogus app chart
    - **It:** shows nothing by default
    - **It:** sets a default
    - **It:** fails to sets a bogus default
    - **It:** unsets a default
      -  **By:** setting default
      -  **By:** unsetting default
  - **Context:** command+ command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments

## `apps_env_test.go`

- **Describe:** apps env, LApplication
    - **It:** is not shown in the environment listing
    - **It:** is retrieved as empty string with show
    - **It:** is not present in the pushed workload
  - **Context:** command completion
    - **It:** matches empty app prefix
    - **It:** matches empty var prefix
    - **It:** does not match unknown app prefix
    - **It:** does not match unknown var prefix
    - **It:** does not match bogus arguments
    - **It:** creates the relevant secret
    - **It:** is shown in the environment listing
  - **Context:** list command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
    - **It:** is retrieved with show
  - **Context:** show command completion
    - **It:** matches empty app prefix
    - **It:** matches empty var prefix
    - **It:** does not match unknown app prefix
    - **It:** does not match unknown var prefix
    - **It:** does not match bogus arguments
    - **It:** is injected into the pushed workload
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
    - **It:** modifies and restarts the app
    - **It:** modifies and restarts the app

## `apps_test.go`

- **Describe:** Apps, LApplication
    - **It:** rejects names not fitting kubernetes requirements
    - **It:** creates the app
  - **Context:** with configuration
    - **It:** creates the app with instance count, configurations, and environment
  - **Context:** manifest
    - **It:** is possible to get a manifest
    - **It:** creates the app with environment variables
    - **It:** creates the workload
    - **It:** rejects mixed origins
    - **It:** rejects a bad provider specification
    - **It:** rejects a bad provider specification for a wrong git url
    - **It:** rejects a bad specification
    - **It:** pushes the app successfully repository alone
      -  **By:** deleting the app
    - **It:** pushes the app successfully repository + branch name
      -  **By:** deleting the app
    - **It:** pushes the app successfully repository + commit id
      -  **By:** deleting the app
    - **It:** rejects without a token
    - **It:** pushes the app when providing a proper token
    - **It:** respects the desired number of instances
    - **It:** respects route changes
    - **It:** respects complete route removal
    - **It:** ignores scheme prefixes in routes
  - **Context:** app charts
    - **It:** fails to change the app chart of the running app
    - **It:** respects the desired app chart
    - **It:** respects environment variable changes
    - **It:** will be staged again, and restarted
    - **It:** will be staged again, and NOT restarted
    - **It:** will be staged again, and NOT restarted
    - **It:** will return an error
    - **It:** wont be staged
      -  **By:** deleting the app
    - **It:** pushes successfully
    - **It:** creates no ingresses
    - **It:** creates an ingress matching the custom route
    - **It:** ignores scheme prefixes in the custom route
    - **It:** uses the custom builder to stage
      -  **By:** Pushing a golang app
      -  **By:** checking if the staging is using custom builder image
      -  **By:** Pushing an app that will fail
      -  **By:** waiting for the old staging job to complete
    - **It:** shows the proper status
    - **It:** succeeds when re-pushing a fix
      -  **By:** fixing the problem and pushing the application again
      -  **By:** Pushing an app that will fail
      -  **By:** waiting for the old staging job to fail
    - **It:** succeeds when re-pushing a fix
      -  **By:** fixing the problem and pushing the application again
    - **It:** pushes the same app again successfully
      -  **By:** pushing the app again
      -  **By:** deleting the app
    - **It:** honours the given instance count
      -  **By:** pushing without instance count
      -  **By:** pushing with 0 instance count
      -  **By:** pushing with an instance count
      -  **By:** pushing again, without an instance count
    - **It:** is using the cache PVC
    - **It:** deletes the cache PVC too
    - **It:** shows the staging logs
      -  **By:** pushing the app
    - **It:** deploys a golang app
      -  **By:** checking for the application resource
      -  **By:** deleting the app
      -  **By:** checking the application resource was removed
    - **It:** deploys an app from the current dir
      -  **By:** pushing the app in the current working directory
      -  **By:** deleting the app
    - **It:** deploys an app from the specified dir
      -  **By:** pushing the app in the specified app directory
      -  **By:** deleting the app
  - **Context:** manifest
    - **It:** deploys an app with the desired options
      -  **By:** providing a manifest
      -  **By:** pushing the app specified in the manifest
      -  **By:** verifying the stored settings
      -  **By:** deleting the app
    - **It:** removes the apps ingress when deleting an app
      -  **By:** deleting the app
    - **It:** should not fail for a max-length application name
      -  **By:** deleting the app
    - **It:** should not fail for an application name with leading digits
      -  **By:** deleting the app
    - **It:** respects the desired number of instances
    - **It:** deletes a batch of applications
  - **Context:** with configuration
    - **It:** pushes an app with bound configurations
    - **It:** unbinds bound configurations when deleting an app, and then deletes the configuration
      -  **By:** deleting the app
  - **Context:** with explicit domain secret
    - **It:** pushes an app
      -  **By:** Pushing an app normally
      -  **By:** Getting the generated secret for the domain
      -  **By:** Constructing a routing secret
      -  **By:** Uploading the new secret for the domain
      -  **By:** Pushing the app again
      -  **By:** Seeing the generated cert gone
      -  **By:** Seeing the generated secret gone
      -  **By:** Seeing the ingress use the new routing secret
  - **Context:** with environment variable
    - **It:** pushes an app
    - **It:** respects the desired number of instances
  - **Context:** with configuration
    - **It:** respects the bound configurations
    - **It:** lists all apps in the namespace
      -  **By:** out
    - **It:** lists all apps in the namespace in JSON format
    - **It:** shows the details of an app
      -  **By:** out
  - **Context:** details customized
    - **It:** shows the details of a customized app
  - **Context:** exporting customized
    - **It:** fails to export on conflict between destinations
    - **It:** fails to export without a destination
    - **It:** fails to export for an unknown registry destination
    - **It:** fails to export an unknown application
    - **It:** exports the details of a customized app
  - **Context:** exporting
    - **It:** exports the details of an app
    - **It:** correctly handles complex quoting when deploying and exporting an app
      -  **By:** make zero-instance app:  + app
      -  **By:** pushed
      -  **By:** delete app
      -  **By:** deleted
    - **It:** lists apps without instances
      -  **By:** list apps
    - **It:** shows the details of an app without instances
      -  **By:** show details
    - **It:** scales up to a running workload
    - **It:** lists all applications belonging to all namespaces
      -  **By:** out
      -  **By:** deploying an app
      -  **By:** getting the current logs in full
      -  **By:** ----------------------------------
      -  **By:** fmt.SprintfLOGS = %d lines raw), logLength)
      -  **By:** fmt.SprintfLOG_ [%3d]: %s, idx, line)
      -  **By:** ----------------------------------
      -  **By:** fmt.SprintfLOGS = %d lines filtered), logLength)
      -  **By:** fmt.SprintfSKIP %d lines, logLength)
      -  **By:** removing the app
    - **It:** shows the staging logs
    - **It:** follows logs
      -  **By:** read all the logs
      -  **By:** get to the end of logs
      -  **By:** ----------------------------------
      -  **By:** fmt.SprintfSCAN [%3d], i)
      -  **By:** fmt.SprintfSKIP [%3d]: %s, i, scanner.Text))
      -  **By:** ----------------------------------
      -  **By:** adding new logs
      -  **By:** checking the latest log
    - **It:** executes a command in the applications container one of the pods
    - **It:** deploys successfully
  - **Context:** command+ command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
- **Describe:** Custom chart-value
  - **Context:** with chart-value:
    - **It:** appListeningPort, pushes an app
  - **Context:** without chart-value:
    - **It:** appListeningPort, pushes an app

## `bindings_test.go`

- **Describe:** Bounds between Apps & Configurations, LApplication
    - **It:** shows the bound app for configurations list, and vice versa

## `configurations_test.go`

- **Describe:** Configurations, LConfiguration
    - **It:** shows all created configurations
    - **It:** lists configurations in JSON format
    - **It:** lists all configurations belonging to all namespaces
    - **It:** creates a configuration
    - **It:** creates an empty configuration
    - **It:** rejects names not fitting kubernetes requirements
    - **It:** fails for missing arguments, not enough, no files
    - **It:** fails for missing arguments, not enough, with files
    - **It:** fails for missing arguments, key without value
    - **It:** fails for a missing path
    - **It:** fails for a directory
    - **It:** deletes a configuration
    - **It:** deletes multiple configurations
    - **It:** doesnt delete a bound configuration
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does match for more than one configuration but only the remaining one
    - **It:** binds a configuration to the application deployment
  - **Context:** command completion
    - **It:** matches empty configuration prefix
    - **It:** matches empty app prefix
    - **It:** does not match unknown configuration prefix
    - **It:** does not match unknown app prefix
    - **It:** does not match bogus arguments
    - **It:** unbinds a configuration from the application deployment
  - **Context:** command completion
    - **It:** matches empty configuration prefix
    - **It:** matches empty app prefix
    - **It:** does not match unknown configuration prefix
    - **It:** does not match unknown app prefix
    - **It:** does not match bogus arguments
    - **It:** shows configuration details
    - **It:** shows a configuration in JSON format
    - **It:** reads from files, and truncates large configuration details
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
    - **It:** it edits the configuration, and restarts the app
  - **Context:** service-owned configurations
      -  **By:** make service instance:  + service
      -  **By:** wait for deployment
      -  **By:** make app:  + appName
      -  **By:** chart:  + chart
      -  **By:** config:  + config
      -  **By:** bind service:  + service
      -  **By:** wait for bound
      -  **By:** done before
    - **It:** doesnt unbind a service-owned configuration
    - **It:** doesnt delete a bound service-owned configuration
    - **It:** doesnt delete any service-owned configuration
      -  **By:** unbind service:  + appName
      -  **By:** wait for unbound
    - **It:** deletes a service-owned configuration after service deletion
      -  **By:** unbind service:  + appName
      -  **By:** wait for unbound
      -  **By:** remove app:  + appName
      -  **By:** remove service instance:  + service
      -  **By:** done after

## `gitconfigs_test.go`

- **Describe:** Gitconfigs, LGitconfig
    - **It:** creates a gitconfig
    - **It:** rejects creating an existing gitconfig
    - **It:** rejects names not fitting kubernetes requirements
    - **It:** rejects unknown git providers
    - **It:** lists gitconfigs
    - **It:** rejects showing an unknown gitconfig
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
  - **Context:** existing gitconfig
    - **It:** shows a gitconfig
    - **It:** deletes a git configuration
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments

## `helpers_test.go`

      -  **By:** Regular login
      -  **By:** Regular login done
      -  **By:** OIDC login
      -  **By:** Background: login
      -  **By:** Background: signal run completion
      -  **By:** Waiting for auth code query
      -  **By:** Piping auth data into command
      -  **By:** Waiting for login completion
      -  **By:** OIDC login done
      -  **By:** Check for empty settings
      -  **By:** Check for namespace ` + namespace + `
      -  **By:** Check for user/pass settings
      -  **By:** Check for token settings

## `info_test.go`

- **Describe:** Info, LMisc
    - **It:** succeeds

## `login_test.go`

- **Describe:** Login, LMisc
    - **It:** succeeds with a valid user
    - **It:** succeeds with an interactively entered valid user [fixed bug]
    - **It:** succeeds with OIDC
    - **It:** performs implied logout of previous oidc login
    - **It:** performs implied logout of previous regular login
    - **It:** fails with a non existing user
    - **It:** clears a bogus current namespace
    - **It:** respects the port when one is present [fixed bug]

## `logout_test.go`

- **Describe:** Logout, LMisc
    - **It:** succeeds after regular login
    - **It:** succeeds after a login with OIDC

## `namespaces_test.go`

- **Describe:** Namespaces, LNamespace
    - **It:** has a default namespace
    - **It:** creates and targets an namespace
      -  **By:** switching namespace back to default
    - **It:** rejects creating an existing namespace
    - **It:** rejects names not fitting kubernetes requirements
    - **It:** lists namespaces
    - **It:** lists namespaces in JSON format
    - **It:** rejects showing an unknown namespace
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
  - **Context:** existing namespace
    - **It:** shows a namespace
    - **It:** shows a namespace in JSON format
    - **It:** deletes an namespace
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
    - **It:** rejects targeting an unknown namespace
  - **Context:** existing namespace
    - **It:** shows a namespace
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments

## `services_test.go`

- **Describe:** Services, LService
    - **It:** lists the standard catalog
    - **It:** lists the catalog details
    - **It:** lists the extended catalog
    - **It:** lists the extended catalog details
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
      -  **By:** create it
    - **It:** shows a service which has credentials
      -  **By:** show it
      -  **By:** create it
    - **It:** shows a service which has no credentials
      -  **By:** show it
  - **Context:** customized
    - **It:** shows the customized elements of a service
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
      -  **By:** create it
    - **It:** lists a service
      -  **By:** show it
      -  **By:** wait for deployment
      -  **By:** fmt.Sprintf%s/%s up, namespace, service)
    - **It:** lists services in JSON format
    - **It:** list all services
      -  **By:** create them in different namespaces
      -  **By:** show it
    - **It:** list only the services in the user namespace
      -  **By:** create them in different namespaces
      -  **By:** show it
      -  **By:** verify service deployment
    - **It:** creates a service waiting for completion
      -  **By:** create it
      -  **By:** show it
      -  **By:** fmt.Sprintf%s/%s up, namespace, service)
    - **It:** creates a service
      -  **By:** create it
      -  **By:** show it
      -  **By:** wait for deployment
      -  **By:** fmt.Sprintf%s/%s up, namespace, service)
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match bogus arguments
    - **It:** deletes a service
  - **Context:** bulk deletion
    - **It:** deletes multiple services
    - **It:** does match for more than one service
    - **It:** does match for more than one service but only the remaining one
      -  **By:** create it
      -  **By:** create app
      -  **By:** bind it
      -  **By:** verify binding
    - **It:** fails to delete a bound service
      -  **By:** delete app
      -  **By:** delete it
    - **It:** unbinds and deletes a bound service when forced
  - **Context:** github epinio#2551
      -  **By:** create it
    - **It:** unbinds and deletes a bound service when forced, and can bind again
      -  **By:** bind second service to changed app
      -  **By:** verify binding
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does match for more than one argument
      -  **By:** create it
      -  **By:** wait for deployment
      -  **By:** create app
      -  **By:** verify unbinding
      -  **By:** delete it
    - **It:** binds the service
      -  **By:** bind it
      -  **By:** verify binding /app
      -  **By:** verify binding /show
      -  **By:** verify binding /list
      -  **By:** verify binding /list-all
  - **Context:** command completion
    - **It:** matches empty service prefix
    - **It:** matches empty app prefix
    - **It:** does not match unknown service prefix
    - **It:** does not match unknown app prefix
    - **It:** does not match bogus arguments
      -  **By:** create it
      -  **By:** create app
      -  **By:** wait for deployment
      -  **By:** bind it
      -  **By:** verify binding
    - **It:** unbinds the service
      -  **By:** verify unbinding
      -  **By:** verify unbinding /show
      -  **By:** verify unbinding /list
      -  **By:** verify unbinding /list-all
  - **Context:** command completion
    - **It:** matches empty service prefix
    - **It:** matches empty app prefix
    - **It:** does not match unknown service prefix
    - **It:** does not match unknown app prefix
    - **It:** does not match bogus arguments
    - **It:** fails to port-forward to an unknown service
    - **It:** port-forward a service with a single listening port
      -  **By:** Forwarding on port  + port
    - **It:** port-forward a service with multiple listening ports
      -  **By:** fmt.SprintfForwarding on port %s and %s, port1, port2)
    - **It:** port-forward a service with multiple listening ports and multiple addresses
      -  **By:** fmt.SprintfForwarding on port %s and %s, port1, port2)
    - **It:** Port-forward succeeds after regular login fails after logout
      -  **By:** port-forward with logged in user, expect success ...
      -  **By:** port-forward with logged out user, expect failure ...
  - **Context:** command completion
    - **It:** matches empty prefix
    - **It:** does not match unknown prefix
    - **It:** does not match for more than one argument
      -  **By:** creating service instance:  + service
    - **It:** it edits the service, and restarts the app
      -  **By:** editing instance:  + service
      -  **By:** checking instance:  + service
      -  **By:** waiting for app to resettle:  + appName

## `settings_test.go`

- **Describe:** Settings, LMisc
    - **It:** changes the settings when disabling colors
    - **It:** changes the settings when enabling colors
    - **It:** shows the settings
    - **It:** shows empty settings
    - **It:** shows the settings with the password in plaintext
    - **It:** regenerates the certificate
    - **It:** stores the password in base64
    - **It:** fails accessing the server
    - **It:** fails accessing the server

## `suite_test.go`

*No test defined!*

## `users_test.go`

- **Describe:** Users, LMisc

# Tests description for acceptance/install

## `scenario1_test.go`

- **Describe:** <Scenario1> GKE, epinio-ca
      -  **By:** Domain:  + domain
      -  **By:** Zone:    + zoneID
    - **It:** Installs with loadbalancer IP, custom domain, and pushes an app
      -  **By:** Checking LoadBalancer IP
      -  **By:** Updating DNS Entries
      -  **By:** Checking that DNS entry is correctly propagated
      -  **By:** Installing Epinio
      -  **By:** Connecting to Epinio
      -  **By:** Checking Epinio info command
      -  **By:** Pushing an app
      -  **By:** Delete an app

## `scenario2_test.go`

- **Describe:** <Scenario2> GKE, Letsencrypt-staging, Zero instance
    - **It:** Installs with letsencrypt-staging cert, custom domain and pushes an app with 0 instances
      -  **By:** Creating letsencrypt issuer
      -  **By:** Checking LoadBalancer IP
      -  **By:** Updating DNS Entries
      -  **By:** Checking that DNS entry is correctly propagated
      -  **By:** Installing Epinio
      -  **By:** Connecting to Epinio
      -  **By:** Checking Epinio info command
      -  **By:** Pushing an app with zero instances
      -  **By:** Pushing an app with zero instances, and not verifying certs

## `scenario3_test.go`

- **Describe:** <Scenario3> RKE, Private CA, Configuration, on External Registry
    - **It:** Installs with private CA and pushes an app with configuration
      -  **By:** Installing MetalLB
      -  **By:** Checking LoadBalancer IP
      -  **By:** Configuring local-path storage
      -  **By:** Creating private CA issuer
      -  **By:** Installing Epinio
      -  **By:** Connecting to Epinio
      -  **By:** Checking Epinio info command
      -  **By:** Creating a configuration and pushing an app
      -  **By:** Delete an app

## `scenario4_test.go`

- **Describe:** <Scenario4> EKS, epinio-ca, on S3 storage
    - **It:** Installs with loadbalancer IP, custom domain and pushes an app with env vars
      -  **By:** Checking LoadBalancer IP
      -  **By:** Updating DNS Entries
      -  **By:** Checking that DNS entry is correctly propagated
      -  **By:** Installing Epinio
      -  **By:** Allow internal HTTP registry on EKS 1.24+
      -  **By:** Connecting to Epinio
      -  **By:** Checking Epinio info command
      -  **By:** Targeting workspace namespace
      -  **By:** Creating the application
      -  **By:** Deploying a database with service
      -  **By:** Bind the database to the app
      -  **By:** Pushing an app with Env vars
      -  **By:** Delete an app

## `scenario5_test.go`

- **Describe:** <Scenario5> Azure, Letsencrypt-staging, External Registry
    - **It:** Installs with letsencrypt-staging cert and pushes an app
      -  **By:** Checking LoadBalancer IP
      -  **By:** Updating DNS Entries
      -  **By:** Checking that DNS entry is correctly propagated
      -  **By:** Installing Epinio
      -  **By:** Connecting to Epinio
      -  **By:** Checking Epinio info command
      -  **By:** Pushing an app
      -  **By:** Delete an app

## `scenario6_test.go`

- **Describe:** <Scenario6> Azure, epinio-ca, External Registry
    - **It:** Installs with loadbalancer IP, custom domain and pushes an app
      -  **By:** Checking LoadBalancer IP
      -  **By:** Updating DNS Entries
      -  **By:** Checking that DNS entry is correctly propagated
      -  **By:** Installing Epinio
      -  **By:** Connecting to Epinio
      -  **By:** Checking Epinio info command
      -  **By:** Pushing an app
      -  **By:** Delete an app

## `suite_test.go`

      -  **By:** Upgrading
      -  **By:** Setup And Checks Before Upgrade
      -  **By:** Versions before upgrade
      -  **By:** Deploy application pre-upgrade
      -  **By:** Route:  + beforeRoute
      -  **By:** Create configuration pre-upgrade
      -  **By:** Create service pre-upgrade
      -  **By:** Create custom catalog entry pre-upgrade
      -  **By:** Upgrading actual
      -  **By:** Checks After Upgrade
      -  **By:** Versions after upgrade
      -  **By:** Checking reachability ...
      -  **By:** Checking configuration existence ...
      -  **By:** Checking usability of old service
      -  **By:** Checking service existence ...
      -  **By:** Create configuration post-upgrade
      -  **By:** Create service post-upgrade
      -  **By:** Create custom catalog entry post-upgrade
      -  **By:** Create application post-upgrade
      -  **By:** Route:  + afterRoute
      -  **By:** Teardown After Upgrade
      -  **By:** Installing and configuring the prerequisites
      -  **By:** Expecting a client binary
      -  **By:** Compiling Epinio
      -  **By:** Creating registry secret
      -  **By:** Installing cert-manager
      -  **By:** Installing ingress controller

# Tests description for acceptance/api/v1

## `appchart_index_test.go`

- **Describe:** ChartList Endpoint, LAppchart
    - **It:** lists the known app charts

## `appchart_match_test.go`

- **Describe:** ChartMatch Endpoints, LAppchart
    - **It:** lists the app chart names matching the prefix, none
    - **It:** lists the app chart names matching the prefix, standard
    - **It:** lists the app chart names matching the prefix, all

## `appchart_show_test.go`

- **Describe:** ChartShow Endpoint, LAppchart
    - **It:** lists the details of the app chart
    - **It:** returns a 404 when the chart does not exist

## `application_batch_delete_test.go`

- **Describe:** AppBatchDelete Endpoint, LApplication
    - **It:** removes the applications and unbinds configurations

## `application_create_test.go`

- **Describe:** AppCreate Endpoint, LApplication
    - **It:** creates the app resource
    - **It:** remembers the chart in the app resource
    - **It:** fails for a name not fitting kubernetes requirements
    - **It:** fails creating the app

## `application_delete_test.go`

- **Describe:** AppDelete Endpoint, LApplication
    - **It:** removes the application, unbinds bound configurations
    - **It:** returns a 404 when the namespace does not exist
    - **It:** returns a 404 when the app does not exist

## `application_deploy_test.go`

- **Describe:** AppDeploy Endpoint, LApplication
      -  **By:** creating application resource first
  - **Context:** with staging
    - **It:** cleans up old S3 objects
      -  **By:** uploading the code
      -  **By:** staging the application
      -  **By:** uploading the code again
      -  **By:** staging the application again
      -  **By:** waiting for the new blob to appear
      -  **By:** deploying the application
      -  **By:** waiting for the old blob to be gone
    - **It:** doesnt delete the S3 object
      -  **By:** uploading the code
      -  **By:** staging the application
      -  **By:** staging the application again
  - **Context:** with non-staging using custom container image
    - **It:** returns a success
    - **It:** the app Ingress matches the specified route
      -  **By:** creating application resource first
    - **It:** should fail the second deployment
  - **Context:** from git repository
    - **It:** rejects a bad provider specification
    - **It:** rejects a mismatched git provider

## `application_exec_test.go`

- **Describe:** AppExec Endpoint, LApplication
    - **It:** runs a command and gets the output back
    - **It:** returns an error
    - **It:** works

## `application_export_test.go`

- **Describe:** AppPart Endpoint, LApplication
    - **It:** retrieves the named application part, values
    - **It:** retrieves the named application part, manifest
      -  **By:** stringbodyBytes)
      -  **By:** expecting
    - **It:** returns a 404 when the namespace does not exist
    - **It:** returns a 404 when the app does not exist
    - **It:** returns a 400 when the part does not exist

## `application_importgit_test.go`

- **Describe:** AppImportGit Endpoint, LApplication
    - **It:** fails for no gitURL
    - **It:** fails for wrong gitURL
    - **It:** fails for wrong git revision
    - **It:** imports the git repo in the blob store without specifying revision
    - **It:** imports the git repo in the blob store from a branch
    - **It:** imports the git repo in the blob store from a revision
    - **It:** imports the git repo in the blob store from a short commit revision
    - **It:** imports the git repo from a tag and has the right branch and commit
    - **It:** imports the git repo from a commit and has the right branch

## `application_index_all_test.go`

- **Describe:** AllApps Endpoints, LApplication
    - **It:** lists all applications belonging to all namespaces
    - **It:** doesnt list applications belonging to non-accessible namespaces

## `application_index_test.go`

- **Describe:** Apps Endpoint, LApplication
    - **It:** lists all applications belonging to the namespace
    - **It:** returns a 404 when the namespace does not exist

## `application_logs_test.go`

- **Describe:** AppLogs Endpoint, LApplication
      -  **By:** read the logs
    - **It:** should send the logs
      -  **By:** checking if the logs are right
    - **It:** should follow logs
      -  **By:** get to the end of logs
      -  **By:** adding more logs
      -  **By:** checking the latest log message

## `application_match_test.go`

- **Describe:** ApplicationMatch Endpoint, LApplication
    - **It:** returns an error
      -  **By:** create it
    - **It:** lists all apps for empty prefix
      -  **By:** querying matches
    - **It:** lists no apps matching the prefix
    - **It:** lists all apps matching the prefix
      -  **By:** querying matches
    - **It:** returns a 401 response

## `application_portforward_test.go`

- **Describe:** AppPortForward Endpoint, LApplication
    - **It:** runs a GET through the opened stream and gets the response back
    - **It:** fails with a 400 bad request
    - **It:** runs a GET through the opened stream and gets the response back

## `application_restart_test.go`

- **Describe:** AppRestart Endpoint, LApplication
    - **It:** restarts the app
    - **It:** returns a 404 when the namespace does not exist
    - **It:** returns a 404 when the app does not exist

## `application_show_test.go`

- **Describe:** AppShow Endpoint, LApplication
    - **It:** lists the application data
    - **It:** returns a 404 when the namespace does not exist
    - **It:** returns a 404 when the app does not exist

## `application_stage_test.go`

- **Describe:** AppStage Endpoint, LApplication
      -  **By:** creating application resource first
      -  **By:** creating the other application resource first
      -  **By:** uploading the code of the other
      -  **By:** uploading the code of itself
    - **It:** fails to stage
    - **It:** stages with the previous blob
      -  **By:** uploading the code
      -  **By:** staging the application
      -  **By:** staging the application again
    - **It:** stages with the previous builder image
      -  **By:** uploading the code
      -  **By:** staging the application
      -  **By:** staging the application again
    - **It:** returns a success for a tarball
      -  **By:** uploading the code
      -  **By:** staging the application
      -  **By:** deploying the staged resource
      -  **By:** waiting for the deployment to complete
      -  **By:** confirming at highlevel
    - **It:** returns a success for a zip archive
      -  **By:** uploading the code
      -  **By:** staging the application
      -  **By:** deploying the staged resource
      -  **By:** waiting for the deployment to complete
      -  **By:** confirming at highlevel

## `application_update_test.go`

- **Describe:** AppUpdate Endpoint, LApplication
    - **It:** updates an application with the desired number of instances
    - **It:** returns BadRequest when instances is a negative number
    - **It:** returns BadRequest when instances is not a number
    - **It:** synchronizes the ingresses of the application with the new routes list
    - **It:** synchronizes the ingresses of the application with a new empty routes list
    - **It:** binds a configuration to an app
    - **It:** unbinds a configuration from an app
    - **It:** fails on non existing configuration bindings and does not touch any existing configuration config

## `application_upload_test.go`

- **Describe:** AppUpload Endpoint, LApplication
    - **It:** returns the app response
    - **It:** returns the app response
    - **It:** returns the app response
    - **It:** returns the app response
    - **It:** returns the app response
    - **It:** returns the app response

## `application_validate_cv_test.go`

- **Describe:** AppValidateCV Endpoint, LApplication
    - **It:** returns error when asking for a non existing app
    - **It:** returns error when validating for a non existing chart
    - **It:** returns ok when there are no chart values to validate
    - **It:** returns ok for good chart values
    - **It:** fails for an unknown field
    - **It:** fails for an unknown field type
    - **It:** fails for an integer field with a bad minimum
    - **It:** fails for an integer field with a bad maximum
    - **It:** fails for a value out of range < min
    - **It:** fails for a value out of range > max
    - **It:** fails for a value out of range not in enum
    - **It:** fails for a non-integer value where integer required
    - **It:** fails for a non-numeric value where numeric required
    - **It:** fails for a non-boolean value where boolean required

## `configuration_match_test.go`

- **Describe:** ConfigurationMatch Endpoint, LConfiguration
    - **It:** returns an error
      -  **By:** create it
    - **It:** lists all configurations for empty prefix
      -  **By:** querying matches
    - **It:** lists no configurations matching the prefix
    - **It:** lists all configurations matching the prefix
      -  **By:** querying matches
    - **It:** returns a 401 response

## `configurations_delete_test.go`

- **Describe:** Configurations API Application Endpoints, LConfiguration
    - **It:** returns 404
    - **It:** deletes multiple configurations
    - **It:** deletes a single configuration using the old style
    - **It:** deletes and unbinds them

## `configurations_mutation_test.go`

- **Describe:** Configurations API Application Endpoints, Mutations, LConfiguration
    - **It:** returns a bad request for a non JSON body
    - **It:** returns a bad request for a non-object JSON body
    - **It:** returns a bad request for JSON object without `name` key
    - **It:** returns a not found when the namespace does not exist
  - **Context:** with conflicting configuration
    - **It:** returns a conflict
    - **It:** creates the configuration
    - **It:** creates configurations without data
    - **It:** returns a bad request for a non JSON body
    - **It:** returns a bad request for a non-object JSON body
    - **It:** returns a not found when the namespace does not exist
    - **It:** returns a not found when the configuration does not exist
  - **Context:** with bound applications
    - **It:** returns bad request
    - **It:** unbinds and removes the configuration, when former is requested
  - **Context:** without bound applications
    - **It:** removes the configuration
    - **It:** returns a bad request for a non JSON body
    - **It:** returns a bad request for a non-object JSON body
    - **It:** returns a bad request for JSON object without `name` key
    - **It:** returns a not found when the namespace does not exist
    - **It:** returns a not found when the application does not exist
  - **Context:** with application
    - **It:** returns a not found when the configuration does not exist
  - **Context:** and already bound
    - **It:** returns a note about being bound
    - **It:** binds the configuration
    - **It:** returns a not found when the namespace does not exist
    - **It:** returns a not found when the application does not exist
  - **Context:** with application
    - **It:** returns a not found when the configuration does not exist
  - **Context:** with configuration
  - **Context:** already bound
    - **It:** unbinds the configuration
    - **It:** returns a ok even when the configuration is not bound idempotency

## `configurations_test.go`

- **Describe:** Configurations API Application Endpoints, LConfiguration
    - **It:** lists all configurations in the namespace
    - **It:** returns a 404 when the namespace does not exist
    - **It:** lists all configurations belonging to all namespaces
    - **It:** doesnt list configurations belonging to non-accessible namespaces
    - **It:** lists the configuration data
    - **It:** returns a 404 when the namespace does not exist
    - **It:** returns a 404 when the configuration does not exist
    - **It:** fails for a name not fitting kubernetes requirements
    - **It:** edits the configuration
    - **It:** returns a 404 when the namespace does not exist
    - **It:** returns a 404 when the configuration does not exist
    - **It:** replace the configuration
    - **It:** returns a 404 when the namespace does not exist
    - **It:** returns a 404 when the configuration does not exist
    - **It:** only restarts the app if the configuration has changed

## `gitproxy_test.go`

- **Describe:** Gitproxy endpoint, LMisc
    - **It:** proxies the request to github
    - **It:** fails for unknown endpoint
    - **It:** fails for invalid JSON
    - **It:** fails for non whitelisted APIs

## `healthcheck_test.go`

- **Describe:** Healthcheck endpoint, LMisc
    - **It:** returns OK HTTP 200) without authentication
    - **It:** doesnt include the epinio server version in a header non-authenticated request

## `helpers_api_test.go`

*No test defined!*

## `helpers_test.go`

*No test defined!*

## `info_test.go`

- **Describe:** Info endpoint, LMisc
    - **It:** includes the default builder image in the response
    - **It:** includes the epinio server version in a header

## `me_test.go`

- **Describe:** Me endpoint, LMisc
    - **It:** returns info about the current user
    - **It:** fails getting the current user

## `namespace_auth_test.go`

- **Describe:** Users Namespace, LNamespace
    - **It:** shows the users namespace
    - **It:** doesnt show the other users namespace
    - **It:** doesnt show the admins namespace
    - **It:** list only the user1 namespace
    - **It:** shows every namespace
    - **It:** list every namespace
    - **It:** shows the users namespace
    - **It:** doesnt show the other users namespace

## `namespace_delete_test.go`

- **Describe:** DELETE /api/v1/namespaces/:namespace, LNamespace
    - **It:** deletes an namespace including apps, configurations and services

## `namespaces_test.go`

- **Describe:** Namespaces API Application Endpoints, LNamespace
  - **Context:** Namespaces
    - **It:** lists all namespaces
    - **It:** returns a 401 response
    - **It:** fails for non JSON body
    - **It:** fails for non-object JSON body
    - **It:** fails for JSON object without name key
    - **It:** fails for a known namespace
      -  **By:** creating the same namespace a second time
    - **It:** fails for a name not fitting kubernetes requirements
    - **It:** fails for a restricted namespace
    - **It:** creates a new namespace
    - **It:** lists the namespace data
    - **It:** lists all namespaces for empty prefix
    - **It:** lists no namespaces matching the prefix
    - **It:** lists all namespaces matching the prefix
    - **It:** returns a 401 response

## `service_bind_test.go`

- **Describe:** ServiceBind Endpoint, LService
    - **It:** returns 404
    - **It:** returns 404
    - **It:** binds the services secrets
    - **It:** binds the services secrets
    - **It:** binds the only services secrets with the defined types

## `service_create_test.go`

- **Describe:** ServiceCreate Endpoint, LService
    - **It:** returns an error
    - **It:** returns an error
    - **It:** returns an error
    - **It:** returns success immediately if waiting
    - **It:** returns success after the deploy is ok if not waiting

## `service_delete_test.go`

- **Describe:** ServiceDelete Endpoint, LService
    - **It:** returns 404
    - **It:** returns 404
    - **It:** returns 404
      -  **By:** locate service secret:  + secName + , for:  + serviceName
      -  **By:** locate service helm release:  + releaseName + , for:  + serviceName
    - **It:** deletes the helm release
      -  **By:** assemble url
      -  **By:** fmt.Sprintfassemble request for %s, endpoint)
      -  **By:** curl request
      -  **By:** read response
      -  **By:** fmt.Sprintfdecode response %s, stringrespBody))
      -  **By:** check status
      -  **By:** check helm release removal:  + releaseName
  - **Context:** multi service deletion
    - **It:** deletes multiple services
      -  **By:** fmt.SprintfDeleting services  %+v, serviceNames)
      -  **By:** fmt.Sprintf__Response: %s, stringbodyBytes))
      -  **By:** fmt.SprintfVerifying deletion %+v, serviceNames)

## `service_list_test.go`

- **Describe:** ServiceList Endpoint, LService
    - **It:** returns a 200 with an empty list
    - **It:** returns an empty list
    - **It:** returns the list with the service
    - **It:** returns an empty list
    - **It:** returns a list with service1 in namespace1
    - **It:** returns a list with service2 in namespace2
    - **It:** returns a list with both services

## `service_match_test.go`

- **Describe:** ServiceMatch Endpoint, LService
    - **It:** returns an error
      -  **By:** create it
      -  **By:** show it
      -  **By:** wait for deployment
      -  **By:** fmt.Sprintf%s/%s up, namespace, serviceName)
    - **It:** lists all services for empty prefix
      -  **By:** querying matches
    - **It:** lists no services matching the prefix
    - **It:** lists all services matching the prefix
      -  **By:** querying matches
    - **It:** returns a 401 response

## `service_portforward_test.go`

- **Describe:** ServicePortForward Endpoint, LService
  - **Context:** With ensured namespace
  - **Context:** With ensured service
    - **It:** tests the port-forward API
      -  **By:** assemble url

## `service_show_test.go`

- **Describe:** ServiceShow Endpoint, LService
    - **It:** returns a 404
    - **It:** returns a 404
    - **It:** returns the service with status Ready
      -  **By:** checking the catalog fields
    - **It:** returns the service with name prefixed with [Missing]

## `service_unbind_test.go`

- **Describe:** ServiceUnbind Endpoint, LService
    - **It:** Unbinds the service

## `services_catalog_match_test.go`

- **Describe:** ServiceCatalog Match Endpoint, LService
    - **It:** lists all catalog entries for empty prefix
    - **It:** lists no catalog entries matching the prefix
    - **It:** lists all catalog entries matching the prefix
    - **It:** returns a 401 response

## `services_catalog_test.go`

- **Describe:** ServiceCatalog Endpoint, LService
    - **It:** lists services from the epinio namespace
    - **It:** doesnt list services from namespaces other than epinio

## `suite_test.go`

*No test defined!*

# Tests description for acceptance/apps

## `rails_test.go`

- **Describe:** RubyOnRails
    - **It:** can deploy Rails

## `suite_test.go`

*No test defined!*

## `wordpress_test.go`

- **Describe:** Wordpress
    - **It:** can deploy Wordpress

# Tests description for acceptance/upgrade

## `suite_test.go`

*No test defined!*

## `upgrade_bound_test.go`

- **Describe:** <Upgrade2> Epinio upgrade with bound app and services
    - **It:** Can upgrade epinio bound to a custom service
      -  **By:** Versions before upgrade
      -  **By:** Creating a service
      -  **By:** Wait for deployment
      -  **By:** Pushing Wordpress App
      -  **By:** Bind it
      -  **By:** Verify binding
      -  **By:** Versions after upgrade
      -  **By:** Restarting app

## `upgrade_test.go`

- **Describe:** <Upgrade1> Epinio upgrade with running app
    - **It:** can upgrade epinio
      -  **By:** Versions before upgrade
      -  **By:** Versions after upgrade
      -  **By:** Checking reachability ...
      -  **By:** Creating a service post-upgrade
      -  **By:** wait for deployment
      -  **By:** Creating an application post-upgrade


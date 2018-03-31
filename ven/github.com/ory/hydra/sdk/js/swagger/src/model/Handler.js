/**
 * Hydra OAuth2 & OpenID Connect Server
 * Please refer to the user guide for in-depth documentation: https://ory.gitbooks.io/hydra/content/   Hydra offers OAuth 2.0 and OpenID Connect Core 1.0 capabilities as a service. Hydra is different, because it works with any existing authentication infrastructure, not just LDAP or SAML. By implementing a consent app (works with any programming language) you build a bridge between Hydra and your authentication infrastructure. Hydra is able to securely manage JSON Web Keys, and has a sophisticated policy-based access control you can use if you want to. Hydra is suitable for green- (new) and brownfield (existing) projects. If you are not familiar with OAuth 2.0 and are working on a greenfield project, we recommend evaluating if OAuth 2.0 really serves your purpose. Knowledge of OAuth 2.0 is imperative in understanding what Hydra does and how it works.   The official repository is located at https://github.com/ory/hydra   ### Important REST API Documentation Notes  The swagger generator used to create this documentation does currently not support example responses. To see request and response payloads click on **\"Show JSON schema\"**: ![Enable JSON Schema on Apiary](https://storage.googleapis.com/ory.am/hydra/json-schema.png)   The API documentation always refers to the latest tagged version of ORY Hydra. For previous API documentations, please refer to https://github.com/ory/hydra/blob/<tag-id>/docs/api.swagger.yaml - for example:  0.9.13: https://github.com/ory/hydra/blob/v0.9.13/docs/api.swagger.yaml 0.8.1: https://github.com/ory/hydra/blob/v0.8.1/docs/api.swagger.yaml
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

;(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(
      [
        'ApiClient',
        'model/Firewall',
        'model/KeyGenerator',
        'model/Manager',
        'model/Writer'
      ],
      factory
    )
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(
      require('../ApiClient'),
      require('./Firewall'),
      require('./KeyGenerator'),
      require('./Manager'),
      require('./Writer')
    )
  } else {
    // Browser globals (root is window)
    if (!root.HydraOAuth2OpenIdConnectServer) {
      root.HydraOAuth2OpenIdConnectServer = {}
    }
    root.HydraOAuth2OpenIdConnectServer.Handler = factory(
      root.HydraOAuth2OpenIdConnectServer.ApiClient,
      root.HydraOAuth2OpenIdConnectServer.Firewall,
      root.HydraOAuth2OpenIdConnectServer.KeyGenerator,
      root.HydraOAuth2OpenIdConnectServer.Manager,
      root.HydraOAuth2OpenIdConnectServer.Writer
    )
  }
})(this, function(ApiClient, Firewall, KeyGenerator, Manager, Writer) {
  'use strict'

  /**
   * The Handler model module.
   * @module model/Handler
   * @version Latest
   */

  /**
   * Constructs a new <code>Handler</code>.
   * @alias module:model/Handler
   * @class
   */
  var exports = function() {
    var _this = this
  }

  /**
   * Constructs a <code>Handler</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Handler} obj Optional instance to populate.
   * @return {module:model/Handler} The populated <code>Handler</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports()

      if (data.hasOwnProperty('Generators')) {
        obj['Generators'] = ApiClient.convertToType(data['Generators'], {
          String: KeyGenerator
        })
      }
      if (data.hasOwnProperty('H')) {
        obj['H'] = Writer.constructFromObject(data['H'])
      }
      if (data.hasOwnProperty('Manager')) {
        obj['Manager'] = Manager.constructFromObject(data['Manager'])
      }
      if (data.hasOwnProperty('ResourcePrefix')) {
        obj['ResourcePrefix'] = ApiClient.convertToType(
          data['ResourcePrefix'],
          'String'
        )
      }
      if (data.hasOwnProperty('W')) {
        obj['W'] = Firewall.constructFromObject(data['W'])
      }
    }
    return obj
  }

  /**
   * @member {Object.<String, module:model/KeyGenerator>} Generators
   */
  exports.prototype['Generators'] = undefined
  /**
   * @member {module:model/Writer} H
   */
  exports.prototype['H'] = undefined
  /**
   * @member {module:model/Manager} Manager
   */
  exports.prototype['Manager'] = undefined
  /**
   * @member {String} ResourcePrefix
   */
  exports.prototype['ResourcePrefix'] = undefined
  /**
   * @member {module:model/Firewall} W
   */
  exports.prototype['W'] = undefined

  return exports
})
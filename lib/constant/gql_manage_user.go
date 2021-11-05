package constant

const CreateUserDocument = `
mutation createUser($userInfo: CreateUserInput!, $params: String, $identity: CreateUserIdentityInput, $keepPassword: Boolean, $resetPasswordOnFirstLogin: Boolean) {
  createUser(userInfo: $userInfo, params: $params, identity: $identity, keepPassword: $keepPassword, resetPasswordOnFirstLogin: $resetPasswordOnFirstLogin) {
    id
    arn
    userPoolId
    status
    username
    email
    emailVerified
    phone
    phoneVerified
    unionid
    openid
    nickname
    registerSource
    photo
    password
    oauth
    token
    tokenExpiredAt
    loginsCount
    lastLogin
    lastIP
    signedUp
    blocked
    isDeleted
    device
    browser
    company
    name
    givenName
    familyName
    middleName
    profile
    preferredUsername
    website
    gender
    birthdate
    zoneinfo
    locale
    address
    formatted
    streetAddress
    locality
    region
    postalCode
    city
    province
    country
    createdAt
    updatedAt
    externalId
  }
}
`
const CreateUserWithCustomDataDocument = `
mutation createUserWithCustomData($userInfo: CreateUserInput!, $keepPassword: Boolean, $params: String) {
  createUser(userInfo: $userInfo, keepPassword: $keepPassword, params: $params) {
    id
    arn
    userPoolId
    status
    username
    email
    emailVerified
    phone
    phoneVerified
    unionid
    openid
    nickname
    registerSource
    photo
    password
    oauth
    token
    tokenExpiredAt
    loginsCount
    lastLogin
    lastIP
    signedUp
    blocked
    isDeleted
    device
    browser
    company
    name
    givenName
    familyName
    middleName
    profile
    preferredUsername
    website
    gender
    birthdate
    zoneinfo
    locale
    address
    formatted
    streetAddress
    locality
    region
    postalCode
    city
    province
    country
    createdAt
    updatedAt
    externalId
    customData {
      key
      value
      dataType
      label
    }
  }
}
`

const UpdateUserDocument = `
mutation updateUser($id: String, $input: UpdateUserInput!) {
  updateUser(id: $id, input: $input) {
    id
    arn
    userPoolId
    status
    username
    email
    emailVerified
    phone
    phoneVerified
    unionid
    openid
    nickname
    registerSource
    photo
    password
    oauth
    token
    tokenExpiredAt
    loginsCount
    lastLogin
    lastIP
    signedUp
    blocked
    isDeleted
    device
    browser
    company
    name
    givenName
    familyName
    middleName
    profile
    preferredUsername
    website
    gender
    birthdate
    zoneinfo
    locale
    address
    formatted
    streetAddress
    locality
    region
    postalCode
    city
    province
    country
    createdAt
    updatedAt
    externalId
  }
}
`

const DeleteUserDocument = `
mutation deleteUser($id: String!) {
  deleteUser(id: $id) {
    message
    code
  }
}
`

const BatchDeleteUserDocument = `
mutation deleteUsers($ids: [String!]!) {
  deleteUsers(ids: $ids) {
    message
    code
  }
}
`

const BatchGetUserDocument = `
query userBatch($ids: [String!]!, $type: String) {
  userBatch(ids: $ids, type: $type) {
    identities {
      openid
      userIdInIdp
      userId
      connectionId
      isSocial
      provider
      type
      userPoolId
    }
    id
    arn
    userPoolId
    status
    username
    email
    emailVerified
    phone
    phoneVerified
    unionid
    openid
    nickname
    registerSource
    photo
    password
    oauth
    token
    tokenExpiredAt
    loginsCount
    lastLogin
    lastIP
    signedUp
    blocked
    isDeleted
    device
    browser
    company
    name
    givenName
    familyName
    middleName
    profile
    preferredUsername
    website
    gender
    birthdate
    zoneinfo
    locale
    address
    formatted
    streetAddress
    locality
    region
    postalCode
    city
    province
    country
    createdAt
    updatedAt
    externalId
  }
}
`

const BatchGetUserWithCustomDocument = `
query userBatchWithCustomData($ids: [String!]!, $type: String) {
  userBatch(ids: $ids, type: $type) {
    identities {
      openid
      userIdInIdp
      userId
      connectionId
      isSocial
      provider
      type
      userPoolId
    }
    id
    arn
    userPoolId
    status
    username
    email
    emailVerified
    phone
    phoneVerified
    unionid
    openid
    nickname
    registerSource
    photo
    password
    oauth
    token
    tokenExpiredAt
    loginsCount
    lastLogin
    lastIP
    signedUp
    blocked
    isDeleted
    device
    browser
    company
    name
    givenName
    familyName
    middleName
    profile
    preferredUsername
    website
    gender
    birthdate
    zoneinfo
    locale
    address
    formatted
    streetAddress
    locality
    region
    postalCode
    city
    province
    country
    createdAt
    updatedAt
    externalId
    customData {
      key
      value
      dataType
      label
    }
  }
}

`

const ListArchivedUsersDocument = `
query archivedUsers($page: Int, $limit: Int) {
  archivedUsers(page: $page, limit: $limit) {
    totalCount
    list {
      id
      arn
      status
      userPoolId
      username
      email
      emailVerified
      phone
      phoneVerified
      unionid
      openid
      nickname
      registerSource
      photo
      password
      oauth
      token
      tokenExpiredAt
      loginsCount
      lastLogin
      lastIP
      signedUp
      blocked
      isDeleted
      device
      browser
      company
      name
      givenName
      familyName
      middleName
      profile
      preferredUsername
      website
      gender
      birthdate
      zoneinfo
      locale
      address
      formatted
      streetAddress
      locality
      region
      postalCode
      city
      province
      country
      createdAt
      updatedAt
      externalId
    }
  }
}
`
const FindUserDocument = `
query findUser($email: String, $phone: String, $username: String, $externalId: String, $identity: FindUserByIdentityInput) {
  findUser(email: $email, phone: $phone, username: $username, externalId: $externalId, identity: $identity) {
    id
    arn
    userPoolId
    status
    username
    email
    emailVerified
    phone
    phoneVerified
    unionid
    openid
    nickname
    registerSource
    photo
    password
    oauth
    token
    tokenExpiredAt
    loginsCount
    lastLogin
    lastIP
    signedUp
    blocked
    isDeleted
    device
    browser
    company
    name
    givenName
    familyName
    middleName
    profile
    preferredUsername
    website
    gender
    birthdate
    zoneinfo
    locale
    address
    formatted
    streetAddress
    locality
    region
    postalCode
    city
    province
    country
    createdAt
    updatedAt
    externalId
  }
}
`

const FindUserWithCustomDocument = `
query findUserWithCustomData($email: String, $phone: String, $username: String, $externalId: String) {
  findUser(email: $email, phone: $phone, username: $username, externalId: $externalId) {
    id
    arn
    userPoolId
    status
    username
    email
    emailVerified
    phone
    phoneVerified
    unionid
    openid
    nickname
    registerSource
    photo
    password
    oauth
    token
    tokenExpiredAt
    loginsCount
    lastLogin
    lastIP
    signedUp
    blocked
    isDeleted
    device
    browser
    company
    name
    givenName
    familyName
    middleName
    profile
    preferredUsername
    website
    gender
    birthdate
    zoneinfo
    locale
    address
    formatted
    streetAddress
    locality
    region
    postalCode
    city
    province
    country
    createdAt
    updatedAt
    externalId
    customData {
      key
      value
      dataType
      label
    }
  }
}

`

const SearchUserDocument = `
query searchUser($query: String!, $fields: [String], $page: Int, $limit: Int, $departmentOpts: [SearchUserDepartmentOpt], $groupOpts: [SearchUserGroupOpt], $roleOpts: [SearchUserRoleOpt]) {
  searchUser(query: $query, fields: $fields, page: $page, limit: $limit, departmentOpts: $departmentOpts, groupOpts: $groupOpts, roleOpts: $roleOpts) {
    totalCount
    list {
      id
      arn
      userPoolId
      status
      username
      email
      emailVerified
      phone
      phoneVerified
      unionid
      openid
      nickname
      registerSource
      photo
      password
      oauth
      token
      tokenExpiredAt
      loginsCount
      lastLogin
      lastIP
      signedUp
      blocked
      isDeleted
      device
      browser
      company
      name
      givenName
      familyName
      middleName
      profile
      preferredUsername
      website
      gender
      birthdate
      zoneinfo
      locale
      address
      formatted
      streetAddress
      locality
      region
      postalCode
      city
      province
      country
      createdAt
      updatedAt
      externalId
    }
  }
}
`

const SearchUserWithCustomDocument = `
query searchUserWithCustomData($query: String!, $fields: [String], $page: Int, $limit: Int, $departmentOpts: [SearchUserDepartmentOpt], $groupOpts: [SearchUserGroupOpt], $roleOpts: [SearchUserRoleOpt]) {
  searchUser(query: $query, fields: $fields, page: $page, limit: $limit, departmentOpts: $departmentOpts, groupOpts: $groupOpts, roleOpts: $roleOpts) {
    totalCount
    list {
      id
      arn
      userPoolId
      status
      username
      email
      emailVerified
      phone
      phoneVerified
      unionid
      openid
      nickname
      registerSource
      photo
      password
      oauth
      token
      tokenExpiredAt
      loginsCount
      lastLogin
      lastIP
      signedUp
      blocked
      isDeleted
      device
      browser
      company
      name
      givenName
      familyName
      middleName
      profile
      preferredUsername
      website
      gender
      birthdate
      zoneinfo
      locale
      address
      formatted
      streetAddress
      locality
      region
      postalCode
      city
      province
      country
      createdAt
      updatedAt
      externalId
      customData {
        key
        value
        dataType
        label
      }
    }
  }
}
`

const RefreshUserTokenDocument = `
mutation refreshToken($id: String) {
  refreshToken(id: $id) {
    token
    iat
    exp
  }
}
`

const GetUserGroupsDocument = `
query getUserGroups($id: String!) {
  user(id: $id) {
    groups {
      totalCount
      list {
        code
        name
        description
        createdAt
        updatedAt
      }
    }
  }
}
`
const AddUserToGroupDocument = `
mutation addUserToGroup($userIds: [String!]!, $code: String) {
  addUserToGroup(userIds: $userIds, code: $code) {
    message
    code
  }
}
`

const RemoveUserInGroupDocument = `
mutation removeUserFromGroup($userIds: [String!]!, $code: String) {
  removeUserFromGroup(userIds: $userIds, code: $code) {
    message
    code
  }
}`

const GetUserRolesDocument = `
query getUserRoles($id: String!, $namespace: String) {
  user(id: $id) {
    roles(namespace: $namespace) {
      totalCount
      list {
        id
        code
        namespace
        arn
        description
        createdAt
        updatedAt
        parent {
          code
          namespace
          arn
          description
          createdAt
          updatedAt
        }
      }
    }
  }
}
`
const AddUserToRoleDocument = `
mutation assignRole($namespace: String, $roleCode: String, $roleCodes: [String], $userIds: [String!], $groupCodes: [String!], $nodeCodes: [String!]) {
  assignRole(namespace: $namespace, roleCode: $roleCode, roleCodes: $roleCodes, userIds: $userIds, groupCodes: $groupCodes, nodeCodes: $nodeCodes) {
    message
    code
  }
}
`

const RemoveUserInRoleDocument = `
mutation revokeRole($namespace: String, $roleCode: String, $roleCodes: [String], $userIds: [String!], $groupCodes: [String!], $nodeCodes: [String!]) {
  revokeRole(namespace: $namespace, roleCode: $roleCode, roleCodes: $roleCodes, userIds: $userIds, groupCodes: $groupCodes, nodeCodes: $nodeCodes) {
    message
    code
  }
}
`

const ListUserAuthorizedResourcesDocument = `
query listUserAuthorizedResources($id: String!, $namespace: String, $resourceType: String) {
  user(id: $id) {
    authorizedResources(namespace: $namespace, resourceType: $resourceType) {
      totalCount
      list {
        code
        type
        actions
      }
    }
  }
}
`

const SetUdvDocument = `
mutation setUdv($targetType: UDFTargetType!, $targetId: String!, $key: String!, $value: String!) {
  setUdv(targetType: $targetType, targetId: $targetId, key: $key, value: $value) {
    key
    dataType
    value
    label
  }
}
`
const SendFirstLoginVerifyEmailDocument = `
mutation sendFirstLoginVerifyEmail($userId: String!, $appId: String!) {
  sendFirstLoginVerifyEmail(userId: $userId, appId: $appId) {
    message
    code
  }
}
`

package constant

const AccessTokenDocument = `query accessToken($userPoolId: String!, $secret: String!) {
  accessToken(userPoolId: $userPoolId, secret: $secret) {
    accessToken
    exp
    iat
  }
}`

const NodeByIdWithMembersDocument = `
    query nodeByIdWithMembers($page: Int, $limit: Int, $sortBy: SortByEnum, $includeChildrenNodes: Boolean, $id: String!) {
  nodeById(id: $id) {
    id
    orgId
    name
    nameI18n
    description
    descriptionI18n
    order
    code
    root
    depth
    createdAt
    updatedAt
    children
    users(page: $page, limit: $limit, sortBy: $sortBy, includeChildrenNodes: $includeChildrenNodes) {
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
}
    `

const UsersDocument = `
    query users($page: Int, $limit: Int, $sortBy: SortByEnum) {
  users(page: $page, limit: $limit, sortBy: $sortBy) {
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

const OrgDocument = `
    query org($id: String!) {
  org(id: $id) {
    id
    rootNode {
      id
      orgId
      name
      nameI18n
      description
      descriptionI18n
      order
      code
      root
      depth
      path
      createdAt
      updatedAt
      children
    }
    nodes {
      id
      orgId
      name
      nameI18n
      description
      descriptionI18n
      order
      code
      root
      depth
      path
      createdAt
      updatedAt
      children
    }
  }
}
    `

const RolesDocument = `
    query roles($namespace: String, $page: Int, $limit: Int, $sortBy: SortByEnum) {
  roles(namespace: $namespace, page: $page, limit: $limit, sortBy: $sortBy) {
    totalCount
    list {
      id
      namespace
      code
      arn
      description
      createdAt
      updatedAt
    }
  }
}
    `

const RoleWithUsersDocument = `
    query roleWithUsers($code: String!, $namespace: String, $page: Int, $limit: Int) {
  role(code: $code, namespace: $namespace) {
    users(page: $page, limit: $limit) {
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
}
    `

const GetUserDepartmentsDocument = `
    query getUserDepartments($id: String!, $orgId: String) {
  user(id: $id) {
    departments(orgId: $orgId) {
      totalCount
      list {
        department {
          id
          orgId
          name
          nameI18n
          description
          descriptionI18n
          order
          code
          root
          depth
          path
          codePath
          namePath
          createdAt
          updatedAt
          children
        }
        isMainDepartment
        joinedAt
      }
    }
  }
}
    `

const LoginByEmailDocument = `
    mutation loginByEmail($input: LoginByEmailInput!) {
  loginByEmail(input: $input) {
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
    `

const LoginByPhoneCodeDocument = `
    mutation loginByPhoneCode($input: LoginByPhoneCodeInput!) {
  loginByPhoneCode(input: $input) {
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

const LoginByPhonePasswordDocument = `
    mutation loginByPhonePassword($input: LoginByPhonePasswordInput!) {
  loginByPhonePassword(input: $input) {
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

const LoginBySubAccountDocument = `
    mutation loginBySubAccount($account: String!, $password: String!, $captchaCode: String, $clientIp: String) {
  loginBySubAccount(account: $account, password: $password, captchaCode: $captchaCode, clientIp: $clientIp) {
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
    `

const LoginByUsernameDocument = `
    mutation loginByUsername($input: LoginByUsernameInput!) {
  loginByUsername(input: $input) {
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

const IsActionAllowedDocument = `
    query isActionAllowed($resource: String!, $action: String!, $userId: String!, $namespace: String) {
  isActionAllowed(resource: $resource, action: $action, userId: $userId, namespace: $namespace)
}
    `

const AllowDocument = `
    mutation allow($resource: String!, $action: String!, $userId: String, $userIds: [String!], $roleCode: String, $roleCodes: [String!], $namespace: String) {
  allow(resource: $resource, action: $action, userId: $userId, userIds: $userIds, roleCode: $roleCode, roleCodes: $roleCodes, namespace: $namespace) {
    message
    code
  }
}
    `

const AuthorizeResourceDocument = `
    mutation authorizeResource($namespace: String, $resource: String, $resourceType: ResourceType, $opts: [AuthorizeResourceOpt]) {
  authorizeResource(namespace: $namespace, resource: $resource, resourceType: $resourceType, opts: $opts) {
    code
    message
  }
}
    `

const GroupsDocument = `
    query groups($userId: String, $page: Int, $limit: Int, $sortBy: SortByEnum) {
  groups(userId: $userId, page: $page, limit: $limit, sortBy: $sortBy) {
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

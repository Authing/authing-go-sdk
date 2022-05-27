package constant

const CreateGroupsDocument = `
mutation createGroup($code: String!, $name: String!, $description: String) {
  createGroup(code: $code, name: $name, description: $description) {
    code
    name
    description
    createdAt
    updatedAt
  }
}
`

const UpdateGroupsDocument = `
mutation updateGroup($code: String!, $name: String, $description: String, $newCode: String) {
  updateGroup(code: $code, name: $name, description: $description, newCode: $newCode) {
    code
    name
    description
    createdAt
    updatedAt
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

const DetailGroupsDocument = `
query group($code: String!) {
  group(code: $code) {
    code
    name
    description
    createdAt
    updatedAt
  }
}
`

const DeleteGroupsDocument = `
mutation deleteGroups($codeList: [String!]!) {
  deleteGroups(codeList: $codeList) {
    message
    code
  }
}
`

const ListGroupsDocument = `
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

const ListGroupUserDocument = `
query groupWithUsers($code: String!, $page: Int, $limit: Int) {
  group(code: $code) {
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

const ListGroupUserWithCustomDocument = `
query groupWithUsersWithCustomData($code: String!, $page: Int, $limit: Int) {
  group(code: $code) {
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
        customData {
          key
          value
          dataType
          label
        }
      }
    }
  }
}

`

const ListGroupAuthorizedResourcesDocument = `
query listGroupAuthorizedResources($code: String!, $namespace: String, $resourceType: String) {
  group(code: $code) {
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

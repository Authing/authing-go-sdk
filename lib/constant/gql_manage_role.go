package constant

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
}`

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
}`

const CreateRoleDocument = `
mutation createRole($namespace: String, $code: String!, $description: String, $parent: String) {
  createRole(namespace: $namespace, code: $code, description: $description, parent: $parent) {
    id
    namespace
    code
    arn
    description
    createdAt
    updatedAt
    parent {
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

const DeleteRoleDocument = `
mutation deleteRole($code: String!, $namespace: String) {
  deleteRole(code: $code, namespace: $namespace) {
    message
    code
  }
}
`

const BatchDeleteRoleDocument = `
mutation deleteRoles($codeList: [String!]!, $namespace: String) {
  deleteRoles(codeList: $codeList, namespace: $namespace) {
    message
    code
  }
}
`

const UpdateRoleDocument = `
mutation updateRole($code: String!, $description: String, $newCode: String, $namespace: String) {
  updateRole(code: $code, description: $description, newCode: $newCode, namespace: $namespace) {
    id
    namespace
    code
    arn
    description
    createdAt
    updatedAt
    parent {
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

const RoleDetailDocument = `
query role($code: String!, $namespace: String) {
  role(code: $code, namespace: $namespace) {
    id
    namespace
    code
    arn
    description
    createdAt
    updatedAt
    parent {
      namespace
      code
      arn
      description
      createdAt
      updatedAt
    }
  }
}`

const AssignRoleDocument = `
mutation assignRole($namespace: String, $roleCode: String, $roleCodes: [String], $userIds: [String!], $groupCodes: [String!], $nodeCodes: [String!]) {
  assignRole(namespace: $namespace, roleCode: $roleCode, roleCodes: $roleCodes, userIds: $userIds, groupCodes: $groupCodes, nodeCodes: $nodeCodes) {
    message
    code
  }
}
`

const RevokeRoleDocument = `
mutation revokeRole($namespace: String, $roleCode: String, $roleCodes: [String], $userIds: [String!], $groupCodes: [String!], $nodeCodes: [String!]) {
  revokeRole(namespace: $namespace, roleCode: $roleCode, roleCodes: $roleCodes, userIds: $userIds, groupCodes: $groupCodes, nodeCodes: $nodeCodes) {
    message
    code
  }
}
`

const ListPoliciesDocument = `
query policyAssignments($namespace: String, $code: String, $targetType: PolicyAssignmentTargetType, $targetIdentifier: String, $page: Int, $limit: Int) {
  policyAssignments(namespace: $namespace, code: $code, targetType: $targetType, targetIdentifier: $targetIdentifier, page: $page, limit: $limit) {
    totalCount
    list {
      code
      targetType
      targetIdentifier
    }
  }
}
`

const AddPoliciesDocument = `
mutation addPolicyAssignments($policies: [String!]!, $targetType: PolicyAssignmentTargetType!, $targetIdentifiers: [String!], $inheritByChildren: Boolean, $namespace: String) {
  addPolicyAssignments(policies: $policies, targetType: $targetType, targetIdentifiers: $targetIdentifiers, inheritByChildren: $inheritByChildren, namespace: $namespace) {
    message
    code
  }
}
`
const RemovePoliciesDocument = `
mutation removePolicyAssignments($policies: [String!]!, $targetType: PolicyAssignmentTargetType!, $targetIdentifiers: [String!], $namespace: String) {
  removePolicyAssignments(policies: $policies, targetType: $targetType, targetIdentifiers: $targetIdentifiers, namespace: $namespace) {
    message
    code
  }
}
`

const ListRoleAuthorizedResourcesDocument = `
query listRoleAuthorizedResources($code: String!, $namespace: String, $resourceType: String) {
  role(code: $code, namespace: $namespace) {
    authorizedResources(resourceType: $resourceType) {
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
const GetRoleUdfValueDocument = `
query udv($targetType: UDFTargetType!, $targetId: String!) {
  udv(targetType: $targetType, targetId: $targetId) {
    key
    dataType
    value
    label
  }
}
`

const BatchGetRoleUdfValueDocument = `
query udfValueBatch($targetType: UDFTargetType!, $targetIds: [String!]!) {
  udfValueBatch(targetType: $targetType, targetIds: $targetIds) {
    targetId
    data {
      key
      dataType
      value
      label
    }
  }
}
`

const SetRoleUdfValueDocument = `
mutation setUdvBatch($targetType: UDFTargetType!, $targetId: String!, $udvList: [UserDefinedDataInput!]) {
  setUdvBatch(targetType: $targetType, targetId: $targetId, udvList: $udvList) {
    key
    dataType
    value
    label
  }
}
`

const BatchSetUdfValueDocument = `
mutation setUdfValueBatch($targetType: UDFTargetType!, $input: [SetUdfValueBatchInput!]!) {
  setUdfValueBatch(targetType: $targetType, input: $input) {
    code
    message
  }
}
`
const RemoveUdfValueDocument = `
mutation removeUdv($targetType: UDFTargetType!, $targetId: String!, $key: String!) {
  removeUdv(targetType: $targetType, targetId: $targetId, key: $key) {
    key
    dataType
    value
    label
  }
}
`

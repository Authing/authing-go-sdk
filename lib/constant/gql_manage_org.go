package constant

const CreateOrgDocument = `
mutation createOrg($name: String!, $code: String, $description: String) {
  createOrg(name: $name, code: $code, description: $description) {
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

const DeleteOrgDocument = `
mutation deleteOrg($id: String!) {
  deleteOrg(id: $id) {
    message
    code
  }
}
`

const ListOrgDocument = `
query orgs($page: Int, $limit: Int, $sortBy: SortByEnum) {
  orgs(page: $page, limit: $limit, sortBy: $sortBy) {
    totalCount
    list {
      id
      rootNode {
        id
        name
        nameI18n
        path
        description
        descriptionI18n
        order
        code
        root
        depth
        createdAt
        updatedAt
        children
      }
      nodes {
        id
        name
        path
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
      }
    }
  }
}

`

const AddOrgNodeDocument = `
mutation addNode($orgId: String!, $parentNodeId: String, $name: String!, $nameI18n: String, $description: String, $descriptionI18n: String, $order: Int, $code: String) {
  addNode(orgId: $orgId, parentNodeId: $parentNodeId, name: $name, nameI18n: $nameI18n, description: $description, descriptionI18n: $descriptionI18n, order: $order, code: $code) {
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

const GetOrgNodeDocument = `
query nodeById($id: String!) {
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
    path
    createdAt
    updatedAt
    children
  }
}

`

const UpdateOrgNodeDocument = `
mutation updateNode($page: Int, $limit: Int, $sortBy: SortByEnum, $includeChildrenNodes: Boolean, $id: String!, $name: String, $code: String, $description: String) {
  updateNode(id: $id, name: $name, code: $code, description: $description) {
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
    users(page: $page, limit: $limit, sortBy: $sortBy, includeChildrenNodes: $includeChildrenNodes) {
      totalCount
    }
  }
}
`

const DeleteOrgNodeDocument = `
mutation deleteNode($orgId: String!, $nodeId: String!) {
  deleteNode(orgId: $orgId, nodeId: $nodeId) {
    message
    code
  }
}
`

const IsRootNodeDocument = `
query isRootNode($nodeId: String!, $orgId: String!) {
  isRootNode(nodeId: $nodeId, orgId: $orgId)
}
`

const MoveNodeDocument = `

mutation moveNode($orgId: String!, $nodeId: String!, $targetParentId: String!) {
  moveNode(orgId: $orgId, nodeId: $nodeId, targetParentId: $targetParentId) {
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
}`

const GetRootNodeDocument = `
query rootNode($orgId: String!) {
  rootNode(orgId: $orgId) {
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
}

`

const AddMembersDocument = `
mutation addMember($page: Int, $limit: Int, $sortBy: SortByEnum, $includeChildrenNodes: Boolean, $nodeId: String, $orgId: String, $nodeCode: String, $userIds: [String!]!, $isLeader: Boolean) {
  addMember(nodeId: $nodeId, orgId: $orgId, nodeCode: $nodeCode, userIds: $userIds, isLeader: $isLeader) {
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
    users(page: $page, limit: $limit, sortBy: $sortBy, includeChildrenNodes: $includeChildrenNodes) {
      totalCount
      list {
        id
        arn
        userPoolId
        username
        status
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

const MoveNodeMembersDocument = `
mutation moveMembers($userIds: [String!]!, $sourceNodeId: String!, $targetNodeId: String!) {
  moveMembers(userIds: $userIds, sourceNodeId: $sourceNodeId, targetNodeId: $targetNodeId) {
    code
    message
  }
}

`

const RemoveNodeMembersDocument = `
mutation removeMember($page: Int, $limit: Int, $sortBy: SortByEnum, $includeChildrenNodes: Boolean, $nodeId: String, $orgId: String, $nodeCode: String, $userIds: [String!]!) {
  removeMember(nodeId: $nodeId, orgId: $orgId, nodeCode: $nodeCode, userIds: $userIds) {
    id
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
      }
    }
  }
}
`

const SetUserMainDepartmentDocument = `
mutation setMainDepartment($userId: String!, $departmentId: String) {
  setMainDepartment(userId: $userId, departmentId: $departmentId) {
    message
    code
  }
}
`

const ListNodeByIdAuthorizedResourcesDocument = `
query listNodeByIdAuthorizedResources($id: String!, $namespace: String, $resourceType: String) {
  nodeById(id: $id) {
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

const ListNodeByCodeAuthorizedResourcesDocument = `
query listNodeByCodeAuthorizedResources($orgId: String!, $code: String!, $namespace: String, $resourceType: String) {
  nodeByCode(orgId: $orgId, code: $code) {
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

const SearchNodesDocument = `
query searchNodes($keyword: String!) {
  searchNodes(keyword: $keyword) {
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
}
`

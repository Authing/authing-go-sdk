package constant

const CreatePolicyDocument = `
mutation createPolicy($namespace: String, $code: String!, $description: String, $statements: [PolicyStatementInput!]!) {
  createPolicy(namespace: $namespace, code: $code, description: $description, statements: $statements) {
    namespace
    code
    isDefault
    description
    statements {
      resource
      actions
      effect
      condition {
        param
        operator
        value
      }
    }
    createdAt
    updatedAt
    assignmentsCount
  }
}

`

const ListPolicyDocument = `
query policies($page: Int, $limit: Int, $namespace: String) {
  policies(page: $page, limit: $limit, namespace: $namespace) {
    totalCount
    list {
      namespace
      code
      description
      createdAt
      updatedAt
      statements {
        resource
        actions
        effect
        condition {
          param
          operator
          value
        }
      }
    }
  }
}
`

const DetailPolicyDocument = `
query policy($namespace: String, $code: String!) {
  policy(code: $code, namespace: $namespace) {
    namespace
    code
    isDefault
    description
    statements {
      resource
      actions
      effect
      condition {
        param
        operator
        value
      }
    }
    createdAt
    updatedAt
  }
}

`

const UpdatePolicyDocument = `
mutation updatePolicy($namespace: String, $code: String!, $description: String, $statements: [PolicyStatementInput!], $newCode: String) {
  updatePolicy(namespace: $namespace, code: $code, description: $description, statements: $statements, newCode: $newCode) {
    namespace
    code
    description
    statements {
      resource
      actions
      effect
      condition {
        param
        operator
        value
      }
    }
    createdAt
    updatedAt
  }
}

`

const DeletePolicyDocument = `
mutation deletePolicy($code: String!, $namespace: String) {
  deletePolicy(code: $code, namespace: $namespace) {
    message
    code
  }
}
`

const BatchDeletePolicyDocument = `
mutation deletePolicies($codeList: [String!]!, $namespace: String) {
  deletePolicies(codeList: $codeList, namespace: $namespace) {
    message
    code
  }
}
`
const PolicyAssignmentsDocument = `
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
const AddAssignmentsDocument = `
mutation addPolicyAssignments($policies: [String!]!, $targetType: PolicyAssignmentTargetType!, $targetIdentifiers: [String!], $inheritByChildren: Boolean, $namespace: String) {
  addPolicyAssignments(policies: $policies, targetType: $targetType, targetIdentifiers: $targetIdentifiers, inheritByChildren: $inheritByChildren, namespace: $namespace) {
    message
    code
  }
}
`
const RemoveAssignmentsDocument = `
mutation removePolicyAssignments($policies: [String!]!, $targetType: PolicyAssignmentTargetType!, $targetIdentifiers: [String!], $namespace: String) {
  removePolicyAssignments(policies: $policies, targetType: $targetType, targetIdentifiers: $targetIdentifiers, namespace: $namespace) {
    message
    code
  }
}
`

const EnablePolicyAssignmentDocument = `
mutation enablePolicyAssignment($policy: String!, $targetType: PolicyAssignmentTargetType!, $targetIdentifier: String!, $namespace: String) {
  enablePolicyAssignment(policy: $policy, targetType: $targetType, targetIdentifier: $targetIdentifier, namespace: $namespace) {
    message
    code
  }
}
`
const DisablePolicyAssignmentDocument = `
mutation disbalePolicyAssignment($policy: String!, $targetType: PolicyAssignmentTargetType!, $targetIdentifier: String!, $namespace: String) {
  disbalePolicyAssignment(policy: $policy, targetType: $targetType, targetIdentifier: $targetIdentifier, namespace: $namespace) {
    message
    code
  }
}

`

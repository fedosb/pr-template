<!--
- Make sure the PR type is selected correctly to properly classify the changes.
- The description should be concise but clear to all team members.
- QA and Post-deploy instructions are important to minimize issues during testing and deployment.
- Check the checklist before submitting the PR to avoid missed items.
-->

## PR Type
<!--
Remove unnecessary items, keep relevant ones.
-->
- [ ] Feature (new functionality)
- [ ] Bugfix (bug fix)
- [ ] Hotfix (critical production bug)
- [ ] Refactoring (no behavior changes)
- [ ] Docs (documentation update)
- [ ] Chore (dependency/infrastructure update)

## Description of Changes
**What was done?**
<!--
Remove unnecessary items, keep relevant ones.
Briefly describe the essence of the changes. Examples:
- Added the `/api/v1/users` endpoint for user management.
- Fixed input validation error in service X.
- Refactored caching module to reduce memory allocations.
-->

## Related Tasks
<!--
- Links to tasks in the tracker (Jira/GitHub Issues):  
  Example: `PROJ-123`
-->

## QA Instructions
<!--
Add instructions for the QA team to verify the changes. For example:
- What scenarios should be tested?
- What data should be used for testing?
-->

## Post-merge / Post-deploy Instructions
<!--
Provide instructions to execute after merge or deployment, if applicable. For example:
- Database migrations.
- Updating configuration files.
-->

## Checklist
- [ ] Test coverage is not reduced (`go test -cover`)
- [ ] Code documentation updated (Swagger, README, etc.)
- [ ] Verified no resource leaks (goroutines, DB connections)
- [ ] Changes are backward compatible (if applicable)

## Additional Notes
<!-- Anything that can be useful for reviewers or testers:
- **Screenshots/logs**:  
  (for UI or error analysis, if applicable)
- **Need to update documentation**:  
  https://company.atlassian.net/wiki/spaces/DEV/sample-page
- **Request example**:
  ```curl
  curl -X POST https://api.example.com/v1/endpoint
  ```
-->
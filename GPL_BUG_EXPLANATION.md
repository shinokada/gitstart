# GPL vs LGPL - The Critical Difference

## 🔴 CRITICAL BUG FOUND

The script was showing "GPLv3" in the menu but selecting "lgpl-3.0" license.

## The Bug

```bash
select license in "MIT" "Apache-2.0" "GPLv3" "None" "Quit"; do
    case "${license}" in
    # ... other cases ...
    GPLv3)
        license_url="lgpl-3.0"  # ← WRONG! This is LGPL, not GPL!
        break
        ;;
```

## Why This Matters

### GPL-3.0 (GNU General Public License v3)
- **Strong copyleft**
- Any derivative work MUST be GPL-3.0
- Cannot link with proprietary code
- Full source code must be disclosed
- Used by: Linux kernel, GCC, Bash

### LGPL-3.0 (Lesser GNU General Public License v3)
- **Weak copyleft**
- Allows linking with proprietary code
- Library modifications must be LGPL
- Applications using the library can be proprietary
- Used by: GNU C Library (glibc), Qt

## Legal Impact

If a user selected "GPLv3" expecting strong copyleft protection:

**What they wanted (GPL-3.0):**
```
User's Code (GPL-3.0)
  ↓
Someone forks it
  ↓
Their fork MUST be GPL-3.0 (strong protection)
```

**What they actually got (LGPL-3.0):**
```
User's Code (LGPL-3.0)
  ↓
Someone forks it or links to it
  ↓
Their code CAN be proprietary (weak protection)
```

## Real-World Scenario

### Scenario 1: Open Source Library
Developer wants to ensure all forks remain open source:
- **Intended:** GPL-3.0 (enforces open source on derivatives)
- **Got:** LGPL-3.0 (allows proprietary derivatives)
- **Result:** Their code could be used in closed-source products

### Scenario 2: Company Internal Tool
Company wants to GPL a project but protect derivatives:
- **Intended:** GPL-3.0 (strong copyleft)
- **Got:** LGPL-3.0 (weak copyleft)
- **Result:** Competitors could use code in proprietary products

## The Fix

```bash
GPLv3)
    license_url="gpl-3.0"  # ✓ CORRECT
    break
    ;;
```

## How to Verify

### Before the fix:
```bash
$ ./gitstart -d test-repo
Select a license:
1) MIT
2) Apache-2.0
3) GPLv3        ← User selects this
4) None
5) Quit
#? 3

$ cat test-repo/LICENSE | head -5
                   GNU LESSER GENERAL PUBLIC LICENSE  ← WRONG!
                       Version 3, 29 June 2007
```

### After the fix:
```bash
$ ./gitstart -d test-repo
Select a license:
1) MIT
2) Apache-2.0
3) GPLv3        ← User selects this
4) None
5) Quit
#? 3

$ cat test-repo/LICENSE | head -5
                    GNU GENERAL PUBLIC LICENSE  ← CORRECT!
                       Version 3, 29 June 2007
```

## License Comparison Table

| Feature | GPL-3.0 | LGPL-3.0 |
|---------|---------|----------|
| **Copyleft Strength** | Strong | Weak |
| **Derivative Works** | Must be GPL | Must be LGPL |
| **Linking** | Must GPL entire work | Can link with proprietary |
| **Commercial Use** | Allowed if GPL | Easier commercial use |
| **Source Code** | Full disclosure | Library mods only |
| **Use Case** | Applications | Libraries |

## Why This Bug is Critical

1. **Legal Compliance:** Wrong license = wrong legal obligations
2. **User Intent:** User explicitly chose GPL, not LGPL
3. **Project Protection:** Affects how derivatives can be used
4. **License Violation:** Could accidentally violate user's intent
5. **Reputation:** Shows attention to detail in licensing

## Related Resources

- [GPL-3.0 Full Text](https://www.gnu.org/licenses/gpl-3.0.en.html)
- [LGPL-3.0 Full Text](https://www.gnu.org/licenses/lgpl-3.0.en.html)
- [GPL vs LGPL Explained](https://www.gnu.org/licenses/gpl-faq.html)
- [FSF License List](https://www.gnu.org/licenses/license-list.html)

## Testing the Fix

```bash
# Manual test
export XDG_CONFIG_HOME=/tmp/test-config
mkdir -p $XDG_CONFIG_HOME/gitstart
echo "testuser" > $XDG_CONFIG_HOME/gitstart/config

./gitstart -d test-gpl-repo
# Select option 3 (GPLv3)
# Check LICENSE file starts with "GNU GENERAL PUBLIC LICENSE" not "LESSER"

grep -i "lesser" test-gpl-repo/LICENSE
# Should return nothing (exit code 1)

grep -i "gnu general public license" test-gpl-repo/LICENSE
# Should find text (exit code 0)
```

## Conclusion

This is NOT a cosmetic bug. Selecting the wrong license has real legal implications and could affect how a project can be used, forked, and distributed. The fix ensures users get exactly what they select.

**Status:** ✅ FIXED (gitstart line 262)

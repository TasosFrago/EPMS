from plan_data import plans
from pydantic import ValidationError

def test_plans(plans):
    for i, plan in enumerate(plans):
        try:
            validated_plan = Plan_t(**plan.dict())
            print(f"Plan {i + 1} is valid: {validated_plan}")
        except ValidationError as e:
            print(f"Plan {i + 1} is invalid: {e}")
            
if __name__ == "__main__":
    test_plans(plans)